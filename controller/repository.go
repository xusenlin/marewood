package controller

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func RepositoryFindAll(c *gin.Context)  {

	var result []models.Repository

	if database.DB.Order("created_at desc").Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": result,
		"msg":"查询成功",
	})

}

//还要克隆过来
func RepositoryCreate(c *gin.Context)  {

	var repository models.Repository

	if err := c.BindJSON(&repository); err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": err.Error(),
			"msg":"出错",
		})
		return
	}

	if database.DB.Create(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库创建出错",
		})
		return
	}
	cmd := exec.Command("git", "clone",repository.Url)
	cmd.Dir = config.Cfg.RepositoryDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": err.Error(),
			"msg":"克隆仓库出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": string(out),
		"msg":"创建成功",
	})
}

//不同时删除仓库目录，因为任务可能还在使用
func RepositoryDestroy(c *gin.Context)  {
	var repository models.Repository

	id := c.Query("id")

	if database.DB.Where("id = ?", id ).Delete(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库删除出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":id,
		"msg":"删除成功",
	})

}

//webHook 更新
func RepositoryUpdate(c *gin.Context)  {}
