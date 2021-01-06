package serviceTask

import (
	"MareWood/config"
	"MareWood/helper"
	"MareWood/models"
	"MareWood/service/serviceRepository"
	"MareWood/sql"
	"fmt"
	"os"
	"strconv"
)

/**
流程
更新代码->检查分支->切换分支->
更新代码->安装依赖->打包->创建目录并复制代码->更新Url->
更新运行次数->更新任务状态->更新仓库状态为空闲->（执行附加脚本，暂不做）->更新终端信息。
*/
func TaskRun(t *models.Task, repository *models.Repository, claims *models.Claims) {

	var terminalOut string
	repositoryId := strconv.Itoa(int(repository.ID))

	out, err := serviceRepository.GitPull(repositoryId)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out

	branch, err := serviceRepository.GetBranch(strconv.Itoa(int(repository.ID)))
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}

	if !helper.InStrArr(t.Branch, branch) {
		taskRunError(t, repository, config.Cfg.AppName+":Compilation failed, the branch does not exist, please try to cut the branch and switch the branch, and then try again!", claims)
		return
	}

	out, err = serviceRepository.GitCheckout(repositoryId, t.Branch)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out
	out, err = serviceRepository.GitPull(repositoryId)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out

	out, err = serviceRepository.InstallDepend(repositoryId, repository.DependTools)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out

	out, err = serviceRepository.RunBuild(repositoryId, t.BuildCommand)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out

	//创建目录并复制代码
	out, err = CopyBuildResultToWebRootDir(strconv.Itoa(int(t.ID)), repositoryId, t.BuildDir)
	if err != nil {
		taskRunError(t, repository, err.Error(), claims)
		return
	}
	terminalOut += out

	sql.DB.Model(&t).
		Update("status", models.TaskStatusSuccess).
		Update("url", config.Cfg.WebsUrl+"/"+strconv.Itoa(int(t.ID))).
		Update("run_quantity", t.RunQuantity+1).
		Update("terminal_info", terminalOut)

	sql.DB.Model(&repository).
		Update("task_status", models.RepoTaskStatusLeisured)

	successMsg := models.Message{
		Type:            models.MsgTypeSuccess,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		NeedNotifySelf:  true,
		Message:         claims.Username + "'s task, " + t.Name + ", has been compiled successfully",
	}
	models.Broadcast <- successMsg

}

func taskRunError(t *models.Task, repository *models.Repository, errOut string, claims *models.Claims) {
	fmt.Println("taskRunError::", errOut)
	sql.DB.Model(&t).Update("status", models.TaskStatusFail).Update("terminal_info", errOut)
	sql.DB.Model(&repository).Update("task_status", models.RepoTaskStatusLeisured)

	errMsg := models.Message{
		Type:            models.MsgTypeError,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		NeedNotifySelf:  true,
		Message:         claims.Username + "'s task, " + t.Name + ", has been compiled failed",
	}
	models.Broadcast <- errMsg

}

func CopyBuildResultToWebRootDir(taskId string, repositoryId string, buildDir string) (string, error) {
	destination := config.Cfg.WebRootDir + "/" + taskId

	if helper.IsDir(destination) {
		if err := helper.RemoveDir(destination); err != nil {
			return "", err
		}
	}

	distDirArg := config.Cfg.RepositoryDir + "/" + repositoryId + "/" + buildDir

	err := os.Rename(distDirArg, destination)

	if err != nil {
		return "", err
	}
	return config.Cfg.AppName + ":Compiled successfully！！ 👏👏👏👏", nil
}
