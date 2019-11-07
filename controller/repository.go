package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"FEDeployService/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

	if database.DB.Order("created_at desc").Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})

}

//还要克隆过来
func RepositoryCreate(c *gin.Context) {

	var repository models.Repository

	if err := c.BindJSON(&repository); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "出错",
		})
		return
	}
	repository.Status = models.RepoStatusProcessing

	if database.DB.Create(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库创建出错",
		})
		return
	}

	go func() {

		out, err := service.GitClone(repository.Url, repository.UserName, repository.Password)
		if err != nil {
			database.DB.Model(&repository).
				Where("id = ?", repository.ID).
				Update("status", models.RepoStatusFail).
				Update("console_output", out)
			return
		}

		database.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("status", models.RepoStatusSuccess).
			Update("console_output", out)

	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "后台正在执行仓库克隆",
		"msg":    "创建成功",
	})
}

//不同时删除仓库目录，因为任务可能还在使用
func RepositoryDestroy(c *gin.Context) {
	var repository models.Repository

	id := c.Query("id")

	if database.DB.Where("id = ?", id).Delete(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库删除出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}

//webHook 更新
func RepositoryUpdate(c *gin.Context) {

	strId := c.Query("id")
	repositoryId, err := strconv.Atoi(strId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "参数出错",
		})
		return
	}

	out, err := service.GitPullByRepositoryId(repositoryId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "git pull error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   out,
		"msg":    "git pull success",
	})
}
