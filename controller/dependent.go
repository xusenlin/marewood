package controller

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DependentSupport(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   config.Cfg.DependTools,
		"msg":    "成功",
	})
}

func ReinstallDepend(c *gin.Context) {

	id := c.Query("id")

	var repository models.Repository

	if database.DB.First(&repository, id).Update("depend_status", models.RepoDependStatusProcessing).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}


	go func() {

		err := serviceRepository.DeleteDepend(repository.Url)
		if err != nil {
			database.DB.Model(&repository).
				Update("depend_status", models.RepoDependStatusFail).
				Update("depend_terminal_info", err.Error())
			return
		}
		out, err := serviceRepository.InstallDepend(repository.Url, repository.DependTools)

		if err != nil {
			database.DB.Model(&repository).
				Update("depend_status", models.RepoDependStatusFail).
				Update("depend_terminal_info", out)
			return
		}
		database.DB.Model(&repository).
			Update("depend_status", models.RepoDependStatusSuccess).
			Update("depend_terminal_info", out)

	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "后台重新安装依赖中",
	})
}
