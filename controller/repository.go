package controller

import (
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

	if database.DB.Order("created_at desc").Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})

}

func RepositoryCreate(c *gin.Context) {

	var repository models.Repository

	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	repository.WebHookSecret = helper.RandSeq(models.RepoWebHookSecretRandSeqLen)
	repository.Status = models.RepoStatusProcessing
	repository.DependStatus = models.RepoDependStatusProcessing

	if database.DB.Create(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	go func() {
		//克隆并更新记录
		out, err := serviceRepository.GitClone(repository.Url, repository.UserName, repository.Password)
		if err != nil {
			database.DB.Model(&repository).
				Where("id = ?", repository.ID).
				Update("status", models.RepoStatusFail).
				Update("depend_status", models.RepoDependStatusFail).
				Update("terminal_info", out)
			return
		}

		database.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("status", models.RepoStatusSuccess).
			Update("terminal_info", out)

		//安装依赖并更新记录
		dependOut, err := serviceRepository.InstallDepend(repository.Url, repository.DependTools)
		if err != nil {
			database.DB.Model(&repository).
				Where("id = ?", repository.ID).
				Update("depend_status", models.RepoDependStatusFail).
				Update("depend_terminal_info", dependOut)
			return
		}

		database.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("depend_status", models.RepoDependStatusSuccess).
			Update("depend_terminal_info", dependOut)

	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建成功,后台正在执行仓库克隆",
	})
}


func RepositoryDestroy(c *gin.Context) {
	var repository models.Repository

	id := c.Query("id")

	if true {//这里后面要根据此仓库是否被使用来确认仓库是否能被删除
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   id,
			"msg":    "无法删除，还有任务在使用此仓库，下面还有未完成代码",
		})
		return
	}

	if database.DB.Where("id = ?", id).Delete(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	//这里删除仓库目录

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}
