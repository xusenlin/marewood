package serviceRepository

import (
	"FEDeployService/config"
	"FEDeployService/helper"
)

func InstallDepend(url string , dependTools string) (string, error) {

	cmd,arg := helper.BuildDependCmd(dependTools)

	return RunCmdOnRepositoryDir(url,cmd,arg)
}

func DeleteDepend(url string) error {

	repositoryName, err := helper.GetRepositoryNameByUrl(url)

	if err != nil {
		return err
	}

	dependDir := config.Cfg.RepositoryDir + "/" + repositoryName + "/node_modules"

	return helper.RemoveDir(dependDir)
}

