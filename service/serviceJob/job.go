package serviceJob

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
)


/**
流程
更新代码->切换分支->
更新代码->安装依赖->打包->创建目录并复制代码->更新Url->
更新运行次数->更新任务状态->更新仓库状态为空闲->（执行附加脚本，暂不做）->更新终端信息。
 */
func JobRun(job *models.Job,repository *models.Repository)  {
	var terminalOut string
	out,err := serviceRepository.GitPullByRepositoryUrl(repository.Url)
	if err != nil{
		jobRunError( job, err.Error())
		return
	}
	terminalOut += out
	out,err  = serviceRepository.GitCheckout(repository.Url, job.Branch)
	if err != nil{
		jobRunError( job, err.Error())
		return
	}
	terminalOut += out
	out,err = serviceRepository.GitPullByRepositoryUrl(repository.Url)
	if err != nil{
		jobRunError( job, err.Error())
		return
	}
	terminalOut += out

	out,err = serviceRepository.InstallDepend(repository.Url,repository.DependTools)
	if err != nil{
		jobRunError( job, err.Error())
		return
	}
	terminalOut += out

	out,err = serviceRepository.RunBuild(repository.Url,job.BuildCommand)
	if err != nil{
		jobRunError( job, err.Error())
		return
	}
	terminalOut += out

	//下面 创建目录并复制代码



}

func jobRunError(job *models.Job,errOut string)  {
	database.DB.Model(&job).Update("status", models.JobStatusFail).Update("terminal_info",errOut)
}

