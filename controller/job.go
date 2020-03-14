package controller

import (
	"FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"FEDeployService/service/serviceJob"
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
	if job.PassWord != "" {
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

	if database.DB.Model(&job).UpdateColumn("web_hook_url", webHookUrl).Error != nil {
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
		Error != nil {
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

func JobUpdateBranch(c *gin.Context) {

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

	if database.DB.Where("id = ?", id).Delete(&models.Job{}).Error != nil {
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

/**
流程
仓库是否忙碌->仓库状态是否正常->加密验证->任务状态进行中->更新仓库状态为繁忙->响应->更新代码->切换分支->
更新代码->安装依赖->打包->创建目录并复制代码->更新Url->
更新运行次数->更新任务状态->更新仓库状态为空闲->（执行附加脚本，暂不做）->更新终端信息。
*/
func JobRun(c *gin.Context) {

	jobId := c.Query("id")
	password := c.Query("password")
	var job models.Job
	var repository models.Repository

	if database.DB.First(&job, jobId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	if database.DB.First(&repository, job.RepositoryId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	if repository.Status != models.RepoStatusSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "仓库已经损坏，请尝试修复对应仓库",
		})
		return
	}
	if repository.Password != "" && repository.Password != helper.DigestString(password) {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "密码不对",
		})
		return
	}
	if repository.JobStatus != models.RepoJobStatusLeisured {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "仓库资源已经被占用",
		})
		return
	}
	//任务状态和仓库的任务状态
	repository.JobStatus = models.RepoJobStatusBusy
	job.Status = models.JobStatusProcessing

	if database.DB.Save(&repository).Save(&job).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	//go
	go serviceJob.JobRun(&job, &repository)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "已经进入后台打包",
		"msg":    "",
	})

}
