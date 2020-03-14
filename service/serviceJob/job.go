package serviceJob

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
	"fmt"
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
	out, err := serviceRepository.GitPullByRepositoryUrl(repository.Url)
	if err != nil {
		jobRunError(job, repository,err.Error())
		return
	}
	terminalOut += out
	out, err = serviceRepository.GitCheckout(repository.Url, job.Branch)
	if err != nil {
		jobRunError(job, repository,err.Error())
		return
	}
	terminalOut += out
	out, err = serviceRepository.GitPullByRepositoryUrl(repository.Url)
	if err != nil {
		jobRunError(job, repository,err.Error())
		return
	}
	terminalOut += out

	out, err = serviceRepository.InstallDepend(repository.Url, repository.DependTools)
	if err != nil {
		jobRunError(job, repository,err.Error())
		return
	}
	terminalOut += out

	out, err = serviceRepository.RunBuild(repository.Url, job.BuildCommand)
	if err != nil {
		jobRunError(job, repository,err.Error())
		return
	}
	terminalOut += out

	//创建目录并复制代码
	out, err = CopyBuildResultToWebRootDir(int(job.ID), repository.Url, job.BuildDir)
	if err != nil {
		jobRunError(job,repository, err.Error())
		return
	}
	terminalOut += out

	database.DB.Model(&job).
		Update("status", models.JobStatusSuccess).
		Update("url", config.Cfg.WebsUrl + "/" + strconv.Itoa(int(job.ID))).
		Update("run_quantity", job.RunQuantity + 1).
		Update("terminal_info", terminalOut)

	database.DB.Model(&repository).
		Update("job_status", models.RepoJobStatusLeisured)

}

func jobRunError(job *models.Job, repository *models.Repository, errOut string) {
	fmt.Println("jobRunError::",errOut)
	database.DB.Model(&job).Update("status", models.JobStatusFail).Update("terminal_info", errOut)
	database.DB.Model(&repository).Update("job_status", models.RepoJobStatusLeisured)

}

func CopyBuildResultToWebRootDir(jobId int, repositoryUrl string, buildDir string) (string, error) {
	destination := config.Cfg.WebRootDir + "/" + strconv.Itoa(jobId)

	if !helper.IsDir(destination) {
		if err := helper.MakeDir(destination); err != nil {
			return "", err
		}
	}

	distDirArg := "./" + buildDir + "/"
	destinationArg := destination + "/"

	return serviceRepository.RunCmdOnRepositoryDir(
		repositoryUrl, "cp", "-rf", distDirArg, destinationArg)
}
