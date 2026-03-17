package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log/slog"
	"marewood/internal/dao"
	"marewood/internal/models"
	pkgErrors "marewood/internal/pkg/errors"
	"marewood/internal/pkg/event"
	"marewood/internal/pkg/jwt"
	"marewood/internal/pkg/logger"
	"os"
	"path/filepath"
	"strings"
)

// RepositoryService 仓库业务逻辑接口
type RepositoryService interface {
	CreateOrUpdate(repo *models.Repository) error
	GetRepository(id uint) (*models.Repository, error)
	DeleteRepository(id uint, claims *jwt.Claims) error
	ListRepositories(page, pageSize int, filters map[string]interface{}) ([]*models.Repository, int64, error)
	ListAll() ([]*models.Repository, error)
	Clone(id uint, claims *jwt.Claims) error
	GitPull(id uint, claims *jwt.Claims) error
	GitCheckoutDot(id uint, claims *jwt.Claims) error
	GetBranches(id uint) ([]string, error)
	GetPackageJson(id uint) (*models.PackageJson, error)
}

type repositoryService struct {
	repoDAO    dao.RepositoryDAO
	gitService GitService
}

// NewRepositoryService 创建仓库服务
func NewRepositoryService(repoDAO dao.RepositoryDAO, gitService GitService) RepositoryService {
	return &repositoryService{
		repoDAO:    repoDAO,
		gitService: gitService,
	}
}

func (s *repositoryService) CreateOrUpdate(repo *models.Repository) error {
	// 检查 URL 是否已存在
	existing, err := s.repoDAO.FindByURL(repo.Url, repo.ID)
	if err != nil {
		return pkgErrors.NewInternalError(err)
	}
	if existing != nil {
		return pkgErrors.NewConflictError("repository url already exists")
	}

	// 设置初始状态
	repo.Status = models.RepositoryStatusProcessing
	repo.TaskStatus = models.RepositoryTaskStatusLeisured
	repo.TerminalInfo = ""

	if repo.ID != 0 {
		return s.repoDAO.Update(repo)
	}
	return s.repoDAO.Create(repo)
}

func (s *repositoryService) GetRepository(id uint) (*models.Repository, error) {
	return s.repoDAO.FindByID(id)
}

func (s *repositoryService) DeleteRepository(id uint, claims *jwt.Claims) error {
	// 查找仓库
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return err
	}

	// 检查是否有关联任务
	taskCount, err := s.repoDAO.CountTasksByRepositoryID(id)
	if err != nil {
		return pkgErrors.NewInternalError(err)
	}
	if taskCount != 0 {
		return pkgErrors.NewValidationError("the repository has tasks, cannot be deleted")
	}

	// 删除数据库记录
	if err := s.repoDAO.Delete(id); err != nil {
		return pkgErrors.NewInternalError(err)
	}

	// 发布事件
	event.RepoSource.PublishMsgExcludeID(event.RepoTypeDestroyOk, &event.RepoData{
		RepositoryId: repo.ID,
		UserId:       claims.ID,
		UserName:     claims.Username,
		Msg:          "destroy success",
	}, claims.ID)

	// 删除代码目录
	return os.RemoveAll(repo.CodeDir())
}

func (s *repositoryService) ListRepositories(page, pageSize int, filters map[string]interface{}) ([]*models.Repository, int64, error) {
	offset := (page - 1) * pageSize
	return s.repoDAO.List(offset, pageSize, filters)
}

func (s *repositoryService) ListAll() ([]*models.Repository, error) {
	repos, err := s.repoDAO.ListAll()
	if err != nil {
		return nil, err
	}
	// 添加提示信息
	for idx, r := range repos {
		name, err := getRepositoryNameByURL(r.Url)
		if err != nil {
			repos[idx].TerminalInfo = "git url err"
		} else {
			repos[idx].TerminalInfo = name
		}
	}

	return repos, nil
}

func (s *repositoryService) Clone(id uint, claims *jwt.Claims) error {
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return err
	}

	// 异步执行克隆
	go s.doClone(repo, claims)

	return nil
}

func (s *repositoryService) doClone(repo *models.Repository, claims *jwt.Claims) {
	repoDir := repo.CodeDir()
	terminalInfo := ""

	// 删除旧目录
	_ = os.RemoveAll(repoDir)
	terminalInfo += "RemoveDir:" + repoDir

	// 执行克隆
	out, err := s.gitService.Clone(repo.UserName, repo.Password, repo.DefaultBranch, repo.Url, repoDir)
	terminalInfo += "\n" + out

	eventType := ""
	eventData := &event.RepoData{
		RepositoryId: repo.ID,
		UserId:       claims.ID,
		UserName:     claims.Username,
	}

	if err == nil {
		repo.Status = models.RepositoryStatusSuccess
		eventType = event.RepoTypeCloneOk
		eventData.Msg = "clone success"
		terminalInfo += "\n🥳🥳🥳🥳 Clone successful  🌼🌼🌼"
	} else {
		repo.Status = models.RepositoryStatusFail
		eventType = event.RepoTypeCloneErr
		eventData.Msg = err.Error()
		terminalInfo += "\n" + err.Error()
	}

	repo.TerminalInfo = terminalInfo
	s.repoDAO.Update(repo)

	event.RepoSource.PublishMsg(eventType, eventData)
}

func (s *repositoryService) GitPull(id uint, claims *jwt.Claims) error {
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return err
	}

	if repo.TaskStatus == models.RepositoryTaskStatusBusy {
		return pkgErrors.NewValidationError("this repository is running other tasks")
	}

	// 执行 git pull
	out, err := s.gitService.Pull(repo.CodeDir())
	terminalInfo := out

	eventData := &event.RepoData{
		RepositoryId: repo.ID,
		UserId:       claims.ID,
		UserName:     claims.Username,
	}

	if err != nil {
		terminalInfo += "\n" + err.Error()
		repo.TerminalInfo = terminalInfo
		s.repoDAO.Update(repo)

		eventData.Msg = err.Error()
		event.RepoSource.PublishMsg(event.RepoTypeEditOk, eventData)
		return pkgErrors.NewInternalError(err)
	}

	terminalInfo += "\n🥳🥳🥳🥳 Git Pull successful 🌼🌼🌼"
	repo.TerminalInfo = terminalInfo
	s.repoDAO.Update(repo)

	eventData.Msg = "git pull success"
	event.RepoSource.PublishMsg(event.RepoTypeEditOk, eventData)

	logger.Slog.Info("git pull", slog.String("out", out))
	return nil
}

func (s *repositoryService) GitCheckoutDot(id uint, claims *jwt.Claims) error {
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return err
	}

	if repo.TaskStatus == models.RepositoryTaskStatusBusy {
		return pkgErrors.NewValidationError("this repository is running other tasks")
	}

	// 执行 git checkout .
	out, err := s.gitService.CheckoutDot(repo.CodeDir())
	terminalInfo := out

	eventData := &event.RepoData{
		RepositoryId: repo.ID,
		UserId:       claims.ID,
		UserName:     claims.Username,
	}

	if err != nil {
		terminalInfo += "\n" + err.Error()
		repo.TerminalInfo = terminalInfo
		s.repoDAO.Update(repo)

		eventData.Msg = err.Error()
		event.RepoSource.PublishMsg(event.RepoTypeEditOk, eventData)
		return pkgErrors.NewInternalError(err)
	}

	terminalInfo += "\n🥳🥳🥳🥳 Git Checkout Dot successful 🌼🌼🌼"
	repo.TerminalInfo = terminalInfo
	s.repoDAO.Update(repo)

	eventData.Msg = "git checkout . success"
	event.RepoSource.PublishMsg(event.RepoTypeEditOk, eventData)

	logger.Slog.Info("git checkout .", slog.String("out", out))
	return nil
}

func (s *repositoryService) GetBranches(id uint) ([]string, error) {
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.gitService.GetBranches(repo.CodeDir())
}

func (s *repositoryService) GetPackageJson(id uint) (*models.PackageJson, error) {
	repo, err := s.repoDAO.FindByID(id)
	if err != nil {
		return nil, err
	}

	packageJsonPath := filepath.Join(repo.CodeDir(), "package.json")
	data, err := ioutil.ReadFile(packageJsonPath)
	if err != nil {
		return nil, pkgErrors.NewValidationError("package.json not found")
	}

	var packageJson models.PackageJson
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return nil, pkgErrors.NewInternalError(err)
	}

	return &packageJson, nil
}

// getRepositoryNameByURL 从 URL 中提取仓库名
func getRepositoryNameByURL(gitURL string) (string, error) {
	if !strings.HasPrefix(gitURL, "http") {
		return "", errors.New("the address must have a http prefix")
	}

	noSuffixURL := strings.TrimSuffix(gitURL, ".git")
	urlArr := strings.Split(noSuffixURL, "/")

	return urlArr[len(urlArr)-1], nil
}
