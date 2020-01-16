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

	repositoryId := c.Query("id")

	//if "push" != c.GetHeader("x-github-event") {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": false,
	//		"data":   "",
	//		"msg":    "webHook不是push事件",
	//	})
	//	return
	//}

	//var signature string
	//
	//signature = helper.GetHeaderSignature(c)
	//bodyContent, err := c.GetRawData()

	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": false,
	//		"data":   "",
	//		"msg":    err.Error(),
	//	})
	//	return
	//}

	var repository models.Repository

	if database.DB.First(&repository, repositoryId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error,
		})
		return
	}

	//因为github he gitlab 不一致，暂时不启用
	//if !helper.VerificationWebHookSecret(repository.WebHookSecret, signature, bodyContent) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": false,
	//		"data":   "",
	//		"msg":    "密钥验证不通过",
	//	})
	//	return
	//}

	go func() {
		out,err := serviceRepository.GitPullAndSaveRecord(repository.Url, repository.ID)

		if err != nil || !helper.HasDependChange(out){//仓库拉取代码失败 或者 没有依赖改变
			return
		}
		dependOut,err := serviceRepository.InstallDepend(repository.Url,repository.DependTools)

		if err != nil {
			repository.DependStatus = models.RepoDependStatusFail
			repository.DependTerminalInfo = dependOut
			database.DB.Save(&repository)
		}

	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "git pull 已经执行",
	})
}
