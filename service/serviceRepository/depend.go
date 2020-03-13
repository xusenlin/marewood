package serviceRepository

import (
	"FEDeployService/config"
	"FEDeployService/helper"
	"errors"
	"os"
)

func InstallDepend(url string , dependTools string) (string, error) {

	cmd,arg := helper.BuildDependCmd(dependTools)

	return runCmdOnRepositoryDir(url,cmd,arg)
}

func DeleteDepend(url string) error {

	repositoryName, err := helper.GetRepositoryNameByUrl(url)
	if err != nil {
		return err
	}

	dependDir := config.Cfg.RepositoryDir + "/" + repositoryName + "/node_modules"

	file, err := os.Stat(dependDir)

	if err != nil || !file.IsDir() {
		return  errors.New("找不到依赖目录或者还没有产生依赖,无法删除依赖")
	}

	removeErr := os.RemoveAll(dependDir)

	if removeErr != nil {
		return removeErr
	}
	return  nil
}

