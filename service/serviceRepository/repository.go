package serviceRepository

import (
	"MareWood/config"
	"MareWood/helper"
	"errors"
	"os"
	"os/exec"
	"strings"
)

//克隆仓库，userName，password可留空
func GitClone(repositoryId string, gitUrl string, userName string, password string) (string, error) {

	var cmd *exec.Cmd

	if userName == "" || password == "" {
		cmd = exec.Command("git", "clone", gitUrl)
	} else {
		authUrl, err := helper.GitUrl2AuthUrl(gitUrl, userName, password)
		if err != nil {
			return "", err
		}
		cmd = exec.Command("git", "clone", authUrl)
	}

	cmd.Dir = config.Cfg.RepositoryDir
	out, err := cmd.CombinedOutput()

	if err != nil {
		return string(out), err
	}
	repositoryName, err := helper.GetRepositoryNameByUrl(gitUrl)

	if err != nil {
		return "", err
	}

	repositoryDir := config.Cfg.RepositoryDir + "/" + repositoryName
	newRepositoryDir := config.Cfg.RepositoryDir + "/" + repositoryId

	err = os.Rename(repositoryDir, newRepositoryDir)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

//通过仓库url更新仓库
func GitPull(repositoryId string) (string, error) {

	return RunCmdOnRepositoryDir(repositoryId, "git", "pull")
}

func DeleteRepository(repositoryId string) error {

	repoDir := config.Cfg.RepositoryDir + "/" + repositoryId

	if !helper.IsDir(repoDir) {
		return nil
	}

	return os.RemoveAll(repoDir)
}

func GetBranch(repositoryId string) ([]string, error) {

	out, err := RunCmdOnRepositoryDir(repositoryId, "git", "branch", "-r")

	if err != nil {
		return []string{}, err
	}

	deleteOrigin := strings.ReplaceAll(string(out), "origin/", "")

	branch := strings.Split(strings.Trim(strings.ReplaceAll(deleteOrigin, " ", ""), "\n"), "\n")

	return branch[1:], nil
}

func GitCheckout(repositoryId string, branch string) (string, error) {

	return RunCmdOnRepositoryDir(repositoryId, "git", "checkout", branch)
}

//仓库URL， 构建命令 test、build、build:dev
func RunBuild(repositoryId string, buildCmd string) (string, error) {

	return RunCmdOnRepositoryDir(repositoryId, "npm", "run", buildCmd)
}

func RunCmdOnRepositoryDir(repositoryId string, cmdName string, arg ...string) (string, error) {


	repositoryDir := config.Cfg.RepositoryDir + "/" + repositoryId

	if !helper.IsDir(repositoryDir) {
		return "", errors.New("找不到仓库目录=>" + repositoryDir)
	}

	cmd := exec.Command(cmdName, arg ...)
	cmd.Dir = repositoryDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(cmdName + "命令意外退出=>\n" + string(out))
	}
	return string(out), nil
}
