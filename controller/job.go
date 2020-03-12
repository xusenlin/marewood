package controller

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func JobFindAll(c *gin.Context) {

	var result []models.Job

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

func JobFindByCategoryId(c *gin.Context) {

	var result []models.Job
	id := c.Query("id")

	if database.DB.Where("category_id = ?", id).Order("created_at desc").Find(&result).Error != nil {
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

func JobCreate(c *gin.Context) {

	var job models.Job

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	job.RunQuantity = 0
	job.Branch = "master"
	job.Status = models.JobStatusLeisured
	if job.PassWord != ""{
		job.PassWord = helper.DigestString(job.PassWord)
	}


	if database.DB.Create(&job).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	webHookUrl := config.Cfg.WebHookUrl + "?id=" + strconv.Itoa(int(job.ID))

	if database.DB.Model(&job).UpdateColumn("web_hook_url",webHookUrl).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	//分类数量
	if database.DB.Model(&models.Category{}).
		Where("id = ?", job.CategoryId).
		UpdateColumn("job_quantity", gorm.Expr("job_quantity + ?", 1)).
		Error != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
	}


	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建任务成功",
	})
}

func JobUpdateBranch(c *gin.Context)  {

	var job models.Job

	id := c.Query("id")
	branch := c.Query("branch")

	if database.DB.Model(&job).
		Where("id = ?", id).
		UpdateColumn("branch", branch).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库更新出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   branch,
		"msg":    "更新成功",
	})
}

func JobDestroy(c *gin.Context) {

	id := c.Query("id")

	if database.DB.Where("id = ?",id).Delete(&models.Job{}).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})
}
