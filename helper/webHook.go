package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

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