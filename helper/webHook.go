package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"strings"
)

//获取请求中的签名
func GetHeaderSignature(c *gin.Context) string {
	signature := c.GetHeader("X-Gitlab-Token") //GitLab

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

//从webHooks push中检查是否有依赖变动
func HasDependChange(out string) (hasChange bool)  {
	if strings.Index(out,"package.json") == -1 {
		return false
	}
	return true
}