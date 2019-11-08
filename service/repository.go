package service

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/models"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
	"os"
	"os/exec"
	"time"
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

func GitPullByRepositoryUrl(url string) (string, error) {

	repositoryName, err := GetRepositoryNameByUrl(url)
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

func VerificationWebHookSecret(webHookSecret string,signature string,bodyContent []byte) bool {

	mac := hmac.New(sha1.New, []byte(webHookSecret))
	mac.Write(bodyContent)

	hash := "sha1=" + hex.EncodeToString(mac.Sum(nil))

	if signature != hash {
		return false
	}

	return  true
}

func RandSeq(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GitPullAndSaveRecord(url string,repositoryId uint) bool {

	out, err := GitPullByRepositoryUrl(url)

	record := models.WebHookRecord{Status:0,TerminalInfo:out,RepositoryId:repositoryId}

	if err != nil {
		record.Status = models.WebHookRecordFail
	}else {
		record.Status = models.WebHookRecordSuccess
	}

	return database.DB.NewRecord(record)
}