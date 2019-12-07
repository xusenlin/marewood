package controller

import (
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//webHook 更新
func RepositoryUpdate(c *gin.Context) {

	raw,_ := c.GetRawData()
	c.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   string(raw),
		"msg":    "GetRawData事件",
	})
	return



	repositoryId := c.Query("id")

	if "push" != c.GetHeader("x-github-event") {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "webHook不是push事件",
		})
		return
	}

	var signature string

	signature = helper.GetHeaderSignature(c)
	bodyContent, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	var repository models.Repository

	if database.DB.First(&repository, repositoryId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error,
		})
		return
	}

	if !helper.VerificationWebHookSecret(repository.WebHookSecret, signature, bodyContent) {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "密钥验证不通过",
		})
		return
	}

	go func() {
		serviceRepository.GitPullAndSaveRecord(repository.Url, repository.ID)
	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "git pull 已经执行",
	})
}
