package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"FEDeployService/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

	isNormal := c.Query("isNormal")

	if isNormal == "1" {
		if database.DB.Order("created_at desc").
			Where("status = ?",models.RepoStatusSuccess).
			Find(&result).Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    database.DB.Error.Error(),
			})
			return
		}
	}else {
		if database.DB.Order("created_at desc").Find(&result).Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    database.DB.Error.Error(),
			})
			return
		}
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

	repository.Status = models.RepoStatusProcessing
	repository.JobStatus = models.RepoJobStatusLeisured
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
				Update("terminal_info", out)
			return
		}

		database.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("status", models.RepoStatusSuccess).
			Update("terminal_info", out)

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
			"msg":    "无法删除，还有任务在使用此仓库",
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
	//删除仓库目录
	err := serviceRepository.DeleteRepository(repository.Url)

	if  err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}

func RepositoryGitPull(c *gin.Context)  {

	var repository models.Repository
	id := c.Query("id")

	if database.DB.First( &repository , id ).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	out,err := serviceRepository.GitPullByRepositoryUrl(repository.Url)

	if  err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   out,
		"msg":    "执行成功",
	})
}

func RepositoryDeleteDepend(c *gin.Context)  {

	var repository models.Repository
	id := c.Query("id")

	if database.DB.First( &repository , id ).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	err := serviceRepository.DeleteDepend(repository.Url)

	if  err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "仓库依赖删除成功，相关任务执行的时候会重新安装",
		"msg":    "执行成功",
	})
}