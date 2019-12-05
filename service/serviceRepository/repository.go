package serviceRepository

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)


//克隆仓库，userName，password可留空
func GitClone(gitUrl string, userName string, password string) (string, error) {

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

	return string(out), nil
}

//通过仓库url更新仓库
func GitPullByRepositoryUrl(url string) (string, error) {

	repositoryName, err := helper.GetRepositoryNameByUrl(url)
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

//验证ebHook Secret
func VerificationWebHookSecret(webHookSecret string,signature string,bodyContent []byte) bool {

	mac := hmac.New(sha1.New, []byte(webHookSecret))
	mac.Write(bodyContent)

	hash := "sha1=" + hex.EncodeToString(mac.Sum(nil))

	if signature != hash {
		return false
	}

	return  true
}

//通过仓库url更新仓库并且保存记录
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

//获取请求中的签名
func GetHeaderSignature(c *gin.Context) string {
	signature := c.GetHeader("HTTP_X_GITLAB_TOKEN") //GitLab

	if "" == signature{
		signature = c.GetHeader("X-Hub-Signature") //Github
	}
	if "" == signature{
		signature = c.GetHeader("password") //Gitee
	}
	return signature
}