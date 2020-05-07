package serviceJob

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
更新代码->切换分支->
更新代码->安装依赖->打包->创建目录并复制代码->更新Url->
更新运行次数->更新任务状态->更新仓库状态为空闲->（执行附加脚本，暂不做）->更新终端信息。
*/
func JobRun(job *models.Job, repository *models.Repository) {

	var terminalOut string
	repositoryId := strconv.Itoa(int(repository.ID))

	out, err := serviceRepository.GitPull(repositoryId)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out
	out, err = serviceRepository.GitCheckout(repositoryId, job.Branch)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out
	out, err = serviceRepository.GitPull(repositoryId)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out

	out, err = serviceRepository.InstallDepend(repositoryId, repository.DependTools)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out

	out, err = serviceRepository.RunBuild(repositoryId, job.BuildCommand)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out

	//创建目录并复制代码
	out, err = CopyBuildResultToWebRootDir(strconv.Itoa(int(job.ID)), repositoryId, job.BuildDir)
	if err != nil {
		jobRunError(job, repository, err.Error())
		return
	}
	terminalOut += out

	sql.DB.Model(&job).
		Update("status", models.JobStatusSuccess).
		Update("url", config.Cfg.WebsUrl+"/"+strconv.Itoa(int(job.ID))).
		Update("run_quantity", job.RunQuantity+1).
		Update("terminal_info", terminalOut)

	sql.DB.Model(&repository).
		Update("job_status", models.RepoJobStatusLeisured)

}

func jobRunError(job *models.Job, repository *models.Repository, errOut string) {
	fmt.Println("jobRunError::", errOut)
	sql.DB.Model(&job).Update("status", models.JobStatusFail).Update("terminal_info", errOut)
	sql.DB.Model(&repository).Update("job_status", models.RepoJobStatusLeisured)

}

func CopyBuildResultToWebRootDir(jobId string, repositoryId string, buildDir string) (string, error) {
	destination := config.Cfg.WebRootDir + "/" + jobId

	if helper.IsDir(destination) {
		if err := helper.RemoveDir(destination); err != nil {
			return "", err
		}
	}

	distDirArg := config.Cfg.RepositoryDir + "/"+ repositoryId + "/" + buildDir

	err := os.Rename(distDirArg, destination)

	if err != nil {
		return "", err
	}
	return config.Cfg.AppName + ":打包成功了！！ 👏👏👏👏", nil
}
