package task

import (
	"errors"
	"github.com/gin-gonic/gin"
	cmd "github.com/xusenlin/command"
	"marewood/conf"
	"marewood/internal/command"
	"marewood/internal/common"
	"marewood/internal/context"
	"marewood/internal/db"
	"marewood/internal/event"
	"marewood/internal/repository"
	"os"
	"path/filepath"
	"time"
)

func Run(c *gin.Context) {
	ctx := context.New(c)
	taskId := c.Query("id")

	if taskId == "" {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}
	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	var task Task

	if err := db.Conn.First(&task, taskId).Error; err != nil {
		ctx.SendErr(err)
		return
	}
	if task.Private && claims.ID != task.CreatedId {
		ctx.SendErr(errors.New("only the owner can run a private repository"))
		return
	}
	var repo repository.Repository
	if err := db.Conn.First(&repo, task.RepositoryId).Error; err != nil {
		ctx.SendErr(err)
		return
	}

	if repo.Status != repository.StatusSuccess {
		ctx.SendErr(errors.New("the status of this repository is not normal"))
		return
	}

	if repo.TaskStatus != repository.TaskStatusLeisured {
		ctx.SendErr(errors.New("this repository is running other tasks"))
		return
	}

	task.Status = StatusProcessing
	repo.TaskStatus = repository.TaskStatusBusy

	if err := db.Conn.Save(&task).Error; err != nil {
		ctx.SendErr(err)
		return
	}
	if err := db.Conn.Save(&repo).Error; err != nil {
		ctx.SendErr(err)
		return
	}

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

	go RunTask(claims, &task, &repo)

	ctx.SendOk("Compiling in the background...", task.ID)
}

func RunTask(claims *common.Claims, task *Task, repo *repository.Repository) {
	var terminalOut string
	var err error

	defer func() {
		if err != nil {
			db.Conn.Model(&task).Update("status", StatusFailed).Update("terminal_info", terminalOut)
			event.TaskSource.PublishMsg(event.TaskTypeBuildFail, &event.TaskData{
				UserId:   claims.ID,
				UserName: claims.Username,
				Msg:      "build fail",
				TaskId:   task.ID,
			})
		}
		db.Conn.Model(&repo).Update("task_status", repository.TaskStatusLeisured)
		event.RepoSource.PublishMsg(event.TaskTypeBuildOk, &event.RepoData{
			UserId:       claims.ID,
			UserName:     claims.Username,
			Msg:          "build end",
			RepositoryId: repo.ID,
		})
	}()

	repoDir := repo.CodeDir()

	terminalOut, err = command.GitPull(repoDir)
	if err != nil {
		return
	}

	err = task.CheckBranch()
	if err != nil {
		terminalOut += "\n😭😭😭CheckBranchError:\n" + err.Error()
		return
	}

	out, err := command.GitCheckout(repoDir, task.Branch)
	if err != nil {
		terminalOut += "\n😭😭😭GitCheckoutError:\n" + out + "\n" + err.Error()
		return
	} else {
		terminalOut += out
	}

	name, arg := BuildDependCmd(repo.DependTools)
	installDepend := cmd.NewTimeoutCmd(name, 5*time.Minute).AddArgs(arg)
	out, err = installDepend.RunInDir(repoDir)
	if err != nil {
		terminalOut += "\n😭😭😭InstallDependError:\n" + out + "\n" + err.Error()
		return
	} else {
		terminalOut += out
	}

	pack := cmd.NewTimeoutCmd("npm", 5*time.Minute).AddArgs("run").AddArgs(task.BuildCommand)
	out, err = pack.RunInDir(repoDir)
	if err != nil {
		terminalOut += "\n😭😭😭Compilation failed:\n" + out + "\n" + err.Error()
		return
	} else {
		terminalOut += out
	}

	dist := filepath.Join(repoDir, task.BuildDir)
	website := filepath.Join(conf.WebRootDir, task.Alias)
	_ = os.RemoveAll(website)
	err = os.Rename(dist, website)
	if err != nil {
		terminalOut += err.Error()
		return
	}
	hash, err := command.GitHash(repoDir)

	if err != nil {
		terminalOut += "\n😭😭😭GitHashError:\n" + hash + "\n" + err.Error()
		return
	}

	terminalOut += "\n🥳🥳🥳🥳  Compilation successful  🌼🌼🌼"

	task.TerminalInfo = terminalOut
	task.RunTotal++
	task.CommitHash = hash
	task.Status = StatusSuccess
	db.Conn.Save(&task)

	event.TaskSource.PublishMsg(event.TaskTypeBuildOk, &event.TaskData{
		UserId:   claims.ID,
		UserName: claims.Username,
		Msg:      "build success",
		TaskId:   task.ID,
	})
}

func BuildDependCmd(tools string) (name string, arg string) {
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
