package serviceRepository

import (
	"MareWood/config"
	"MareWood/helper"
)

func InstallDepend(repositoryId string, dependTools string) (string, error) {

	cmd, arg := helper.BuildDependCmd(dependTools)

	return RunCmdOnRepositoryDir(repositoryId, cmd, arg)
}

func DeleteDepend(repositoryId string) error {

	dependDir := config.Cfg.RepositoryDir + "/" + repositoryId + "/node_modules"

	return helper.RemoveDir(dependDir)
}
