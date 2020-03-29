package controller

import (
	"MareWood/helper"
	"MareWood/models"
	"MareWood/service/serviceJob"
	"MareWood/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JobFindAll(c *gin.Context) {

	result,err := new(models.Job).FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
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

	id := c.Query("id")

	result,err := new(models.Job).FindByCategoryId(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
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
	err := c.ShouldBindJSON(&job)

	if  err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err = job.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}


	err = new(models.Category).CategoryJobQuantityIncrement(job.CategoryId)
	//分类数量
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建任务成功",
	})
}

func JobUpdateBranch(c *gin.Context) {

	id := c.Query("id")
	branch := c.Query("branch")

	err := new(models.Job).UpdateBranch(id,branch)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
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

	err := new(models.Job).Destroy(id)
	if err != nil {
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

	if sql.DB.First(&job, jobId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if sql.DB.First(&repository, job.RepositoryId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
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
	if job.Password != "" && job.Password != helper.DigestString(password) {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "任务密码不对",
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

	if sql.DB.Save(&repository).Save(&job).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
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
