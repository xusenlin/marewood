package service

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/models"
	"errors"
	"os"
	"os/exec"
)

func GitClone(gitUrl string, userName string, password string) (string, error) {

	var cmd *exec.Cmd

	if userName == "" || password == "" {
		cmd = exec.Command("git", "clone", gitUrl)
	} else {
		authUrl, err := GitUrl2AuthUrl(gitUrl, userName, password)
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

	return string(out), nil
}

func GitPullByRepositoryId(repositoryId int) (string, error) {

	var repository models.Repository

	if database.DB.First(&repository, repositoryId).Error != nil {
		return "", database.DB.Error
	}

	repositoryName, err := GetRepositoryNameByUrl(repository.Url)
	if err != nil {
		return "", err
	}

	repositoryDir := config.Cfg.RepositoryDir + "/" + repositoryName

	file, err := os.Stat(repositoryDir)

	if err != nil || !file.IsDir() {
		return "", errors.New("找不到仓库目录=>" + repositoryDir)
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = repositoryDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
