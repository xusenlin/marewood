package services

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/gorm"

	"marewood/conf"
	"marewood/internal/dao"
	"marewood/internal/models"
	pkgErrors "marewood/internal/pkg/errors"
	"marewood/internal/pkg/event"
	"marewood/internal/pkg/jwt"
	"marewood/internal/pkg/logger"

	cmd "github.com/xusenlin/command"
)

// TaskService 任务业务逻辑接口
type TaskService interface {
	CreateOrUpdate(task *models.Task, userID uint) error
	GetTask(id uint) (*models.Task, error)
	DeleteTask(id uint, claims *jwt.Claims) error
	ListTasks(page, pageSize int, filters map[string]interface{}, userID uint, userRole int) ([]*models.Task, int64, error)
	GetAllTags() ([]models.Tag, error)
	UpdateBranch(taskID, repositoryID uint) error
	RunTask(taskID uint, claims *jwt.Claims) error
	ArchiveTask(taskID uint, format string) ([]byte, string, error)
}

type taskService struct {
	taskDAO    dao.TaskDAO
	repoDAO    dao.RepositoryDAO
	historyDAO dao.HistoryDao
	gitService GitService
}

// NewTaskService 创建任务服务
func NewTaskService(taskDAO dao.TaskDAO, repoDAO dao.RepositoryDAO, historyDAO dao.HistoryDao, gitService GitService) TaskService {
	return &taskService{
		taskDAO:    taskDAO,
		repoDAO:    repoDAO,
		historyDAO: historyDAO,
		gitService: gitService,
	}
}

func (s *taskService) CreateOrUpdate(task *models.Task, userID uint) error {
	// 检查 Alias 是否已存在
	existing, err := s.taskDAO.FindByAlias(task.Alias, task.ID)
	if err != nil {
		return pkgErrors.NewInternalError(err)
	}
	if existing != nil {
		return pkgErrors.NewConflictError("alias url already exists")
	}

	// 设置初始状态
	task.Status = models.TaskStatusDefault
	task.CreatedId = userID
	task.CommitHash = "None"

	if task.ID != 0 {
		return s.taskDAO.Update(task)
	}
	return s.taskDAO.Create(task)
}

func (s *taskService) GetTask(id uint) (*models.Task, error) {
	return s.taskDAO.FindByID(id)
}

func (s *taskService) DeleteTask(id uint, claims *jwt.Claims) error {
	// 查找任务
	task, err := s.taskDAO.FindByID(id)
	if err != nil {
		return err
	}

	// 删除任务
	if err := s.taskDAO.Delete(id); err != nil {
		return pkgErrors.NewInternalError(err)
	}

	// 发布事件
	event.TaskSource.PublishMsgExcludeID(event.TaskTypeDestroyOk, &event.TaskData{
		TaskId:   task.ID,
		UserId:   claims.ID,
		UserName: claims.Username,
		Msg:      "destroy success",
	}, claims.ID)

	return nil
}

func (s *taskService) ListTasks(page, pageSize int, filters map[string]interface{}, userID uint, userRole int) ([]*models.Task, int64, error) {
	offset := (page - 1) * pageSize
	isAdmin := userRole >= models.UserRoleAdmin
	return s.taskDAO.List(offset, pageSize, filters, userID, isAdmin)
}

func (s *taskService) GetAllTags() ([]models.Tag, error) {
	return s.taskDAO.GetAllTags()
}

func (s *taskService) UpdateBranch(taskID, repositoryID uint) error {
	// 获取任务
	task, err := s.taskDAO.FindByID(taskID)
	if err != nil {
		return err
	}

	// 获取仓库代码目录
	repo, err := s.repoDAO.FindByID(repositoryID)
	if err != nil {
		return err
	}

	// 检查分支是否存在
	branches, err := s.gitService.GetBranches(repo.CodeDir())
	if err != nil {
		return pkgErrors.NewInternalError(err)
	}

	found := false
	for _, branch := range branches {
		if branch == task.Branch {
			found = true
			break
		}
	}

	if !found {
		return pkgErrors.NewValidationError("the branch does not exist")
	}

	return nil
}

func (s *taskService) RunTask(taskID uint, claims *jwt.Claims) error {
	// 获取任务
	task, err := s.taskDAO.FindByID(taskID)
	if err != nil {
		return err
	}

	// 检查权限
	if task.Private && claims.ID != task.CreatedId {
		return pkgErrors.NewForbiddenError("only the owner can run a private task")
	}

	// 获取仓库
	repo, err := s.repoDAO.FindByID(task.RepositoryId)
	if err != nil {
		return err
	}

	// 检查仓库状态
	if repo.Status != models.RepositoryStatusSuccess {
		return pkgErrors.NewValidationError("the status of this repository is not normal")
	}

	if repo.TaskStatus != models.RepositoryTaskStatusLeisured {
		return pkgErrors.NewValidationError("this repository is running other tasks")
	}

	// 更新状态
	task.Status = models.TaskStatusProcessing
	repo.TaskStatus = models.RepositoryTaskStatusBusy

	if err := s.taskDAO.Update(task); err != nil {
		return pkgErrors.NewInternalError(err)
	}
	if err := s.repoDAO.Update(repo); err != nil {
		return pkgErrors.NewInternalError(err)
	}

	// 发布事件
	event.TaskSource.PublishMsgExcludeID(event.TaskTypeRunOk, &event.TaskData{
		UserId:   claims.ID,
		UserName: claims.Username,
		Msg:      "run success",
		TaskId:   task.ID,
	}, claims.ID)
	event.RepoSource.PublishMsgExcludeID(event.TaskTypeRunOk, &event.RepoData{
		UserId:       claims.ID,
		UserName:     claims.Username,
		Msg:          "run success",
		RepositoryId: repo.ID,
	}, claims.ID)

	logger.Slog.Info("run task", slog.Any("userId", claims.ID), slog.Any("taskId", task.ID), slog.Any("repoId", repo.ID))

	// 异步执行任务
	go s.executeTask(claims, task, repo)

	return nil
}

func (s *taskService) executeTask(claims *jwt.Claims, task *models.Task, repo *models.Repository) {
	var terminalOut string
	var err error

	defer func() {
		if err != nil {
			terminalOut = "😭😭😭RunTaskError:\n" + err.Error()
			task.Status = models.TaskStatusFailed
			task.TerminalInfo = terminalOut
			s.taskDAO.Update(task)

			event.TaskSource.PublishMsg(event.TaskTypeBuildFail, &event.TaskData{
				UserId:   claims.ID,
				UserName: claims.Username,
				Msg:      "build fail",
				TaskId:   task.ID,
			})
			logger.Slog.Error("run task error", slog.Any("userId", claims.ID), slog.Any("taskId", task.ID), slog.Any("repoId", repo.ID), slog.String("err", err.Error()))
		}

		repo.TaskStatus = models.RepositoryTaskStatusLeisured
		s.repoDAO.Update(repo)

		event.RepoSource.PublishMsg(event.TaskTypeBuildOk, &event.RepoData{
			UserId:       claims.ID,
			UserName:     claims.Username,
			Msg:          "build end",
			RepositoryId: repo.ID,
		})
	}()

	repoDir := repo.CodeDir()

	// Git Pull
	terminalOut, err = s.gitService.Pull(repoDir)
	if err != nil {
		return
	}

	// 检查分支
	branches, err := s.gitService.GetBranches(repoDir)
	if err != nil {
		terminalOut += "\n😭😭😭CheckBranchError:\n" + err.Error()
		return
	}

	branchExists := false
	for _, branch := range branches {
		if branch == task.Branch {
			branchExists = true
			break
		}
	}

	if !branchExists {
		terminalOut += "\n😭😭😭CheckBranchError:\nthe branch does not exist"
		err = pkgErrors.NewValidationError("the branch does not exist")
		return
	}

	// Git Checkout
	out, checkoutErr := s.gitCheckout(repoDir, task.Branch)
	if checkoutErr != nil {
		terminalOut += "\n😭😭😭GitCheckoutError:\n" + out + "\n" + checkoutErr.Error()
		err = checkoutErr
		return
	}
	terminalOut += out

	// 安装依赖
	name, arg := buildDependCmd(repo.DependTools)
	installDepend := cmd.NewTimeoutCmd(name, 10*time.Minute).AddArgs(arg)
	out, installErr := installDepend.RunInDir(repoDir)
	if installErr != nil {
		terminalOut += "\n😭😭😭InstallDependError:\n" + out + "\n" + installErr.Error()
		err = installErr
		return
	}
	terminalOut += out

	// 执行构建命令
	pack := cmd.NewTimeoutCmd("npm", 5*time.Minute).AddArgs("run").AddArgs(task.BuildCommand)
	out, buildErr := pack.RunInDir(repoDir)
	if buildErr != nil {
		terminalOut += "\n😭😭😭Compilation failed:\n" + out + "\n" + buildErr.Error()
		err = buildErr
		return
	}
	terminalOut += out

	// 移动构建文件
	dist := filepath.Join(repoDir, task.BuildDir)
	website := task.WebDir()
	_ = os.RemoveAll(website)
	err = os.Rename(dist, website)
	if err != nil {
		terminalOut += err.Error()
		return
	}

	// 获取 commit hash
	hash, hashErr := s.gitService.GetCommitHash(repoDir)
	if hashErr != nil {
		terminalOut += "\n😭😭😭GitHashError:\n" + hash + "\n" + hashErr.Error()
		err = hashErr
		return
	}

	// 创建或更新历史记录
	existingHistory, err := s.historyDAO.FindByTaskIDAndGitHash(task.ID, hash)
	if err != nil && err != gorm.ErrRecordNotFound {
		terminalOut += "\n😭😭😭FindHistoryError:\n" + err.Error()
		err = err
		return
	}

	if existingHistory.ID != 0 {
		existingHistory.UserID = claims.ID
		existingHistory.UserName = claims.Username
		if err := s.historyDAO.Update(existingHistory); err != nil {
			terminalOut += "\n😭😭😭UpdateHistoryError:\n" + err.Error()
			err = err
			return
		}
	} else {
		history := &models.History{
			TaskID:   task.ID,
			Name:     task.Name,
			GitHash:  hash,
			UserID:   claims.ID,
			UserName: claims.Username,
		}
		if err := s.historyDAO.Create(history); err != nil {
			terminalOut += "\n😭😭😭CreateHistoryError:\n" + err.Error()
			err = err
			return
		}
	}

	// 复制文件到历史目录
	historyDir := filepath.Join(conf.HistoryDir, strconv.Itoa(int(task.ID)), hash)
	if err := s.copy(website, historyDir); err != nil {
		terminalOut += "\n😭😭😭CopyHistoryError:\n" + err.Error()
		err = err
		return
	}

	terminalOut += "\n🥳🥳🥳🥳  Compilation successful  🌼🌼🌼"

	// 更新任务状态
	task.TerminalInfo = terminalOut
	task.RunTotal++
	task.CommitHash = hash
	task.Status = models.TaskStatusSuccess
	s.taskDAO.Update(task)

	// 发布成功事件
	event.TaskSource.PublishMsg(event.TaskTypeBuildOk, &event.TaskData{
		UserId:   claims.ID,
		UserName: claims.Username,
		Msg:      "build success",
		TaskId:   task.ID,
	})
}

func (s *taskService) ArchiveTask(taskID uint, format string) ([]byte, string, error) {
	task, err := s.taskDAO.FindByID(taskID)
	if err != nil {
		return nil, "", err
	}

	var archiveFile []byte
	switch format {
	case "zip":
		archiveFile, err = s.zipTask(task)
	case "tar":
		archiveFile, err = s.tarTask(task)
	default:
		return nil, "", fmt.Errorf("unsupported archive format: %s", format)
	}

	filename := fmt.Sprintf("%s_%s_%s.%s", task.Tag, task.Name, task.CommitHash, format)
	return archiveFile, filename, nil
}

func (s *taskService) tarTask(task *models.Task) ([]byte, error) {
	dst := task.WebDir()
	folder, err := os.Open(dst)
	if err != nil {
		return nil, err
	}
	defer folder.Close()

	tarFile := new(bytes.Buffer)
	tw := tar.NewWriter(tarFile)

	err = filepath.Walk(dst, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(dst, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(rel)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return tarFile.Bytes(), err
	}
	if err := tw.Close(); err != nil {
		return tarFile.Bytes(), err
	}

	return tarFile.Bytes(), nil
}

func (s *taskService) zipTask(task *models.Task) ([]byte, error) {
	dst := task.WebDir()
	zipFile := new(bytes.Buffer)

	zw := zip.NewWriter(zipFile)

	err := filepath.Walk(dst, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(dst, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(rel)

		w, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			_, err = io.Copy(w, data)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = zw.Close()
	if err != nil {
		return nil, err
	}

	return zipFile.Bytes(), nil
}

// buildDependCmd 构建依赖安装命令
func buildDependCmd(tools string) (name string, arg string) {
	switch tools {
	case "npm":
		return "npm", "install"
	case "yarn":
		return "yarn", "install"
	case "pnpm":
		return "pnpm", "install"
	default:
		return "npm", "install"
	}
}

// gitCheckout 执行 git checkout
func (s *taskService) gitCheckout(dir, branch string) (string, error) {
	checkout := cmd.New("git").AddArgs("checkout", branch)
	return checkout.RunInDir(dir)
}

func (s *taskService) copy(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(dstPath, data, info.Mode())
	})
}
