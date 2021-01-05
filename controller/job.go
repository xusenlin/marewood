package controller

import (
	"MareWood/models"
	"MareWood/service/serviceJob"
	"MareWood/service/serviceUser"
	"MareWood/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JobFindAll(c *gin.Context) {

	result, err := new(models.Job).FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "Database query error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "search successful",
	})

}

func JobFind(c *gin.Context) {

	name := c.Query("name")

	categoryId, err := strconv.Atoi(c.Query("categoryId"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 8
	}
	var result models.JobPageResult
	if strings.HasPrefix(name, "id:") {
		result, err = models.FindJob("", pageNum, pageSize,
			map[string]interface{}{"id": strings.TrimLeft(name, "id:"),"category_id": categoryId})
	} else {
		result, err = models.FindJob(name, pageNum, pageSize, map[string]interface{}{"category_id": categoryId})
	}

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
		"data":   result,
		"msg":    "search successful",
	})

}

func JobCreate(c *gin.Context) {

	var job models.Job
	err := c.ShouldBindJSON(&job)

	if err != nil {
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
		"msg":    "Successfully created task",
	})
}

func JobUpdateBranch(c *gin.Context) {

	id := c.Query("id")
	branch := c.Query("branch")

	var job models.Job

	if sql.DB.First(&job, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if job.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
		})
		return
	}
	if err := job.UpdateBranch(branch); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
			"msg":    "Database update error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   branch,
		"msg":    "update completed",
	})
}

func JobDestroy(c *gin.Context) {

	id := c.Query("id")
	var job models.Job

	if sql.DB.First(&job, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if job.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
		})
		return
	}

	err := job.Destroy(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	claims, _ := serviceUser.GetJwtClaimsByContext(c)
	msg := models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "”deleted the task “" + job.Name + "”",
	}
	models.Broadcast <- msg
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
	})
}

func JobUpdateField(c *gin.Context) {

	id := c.Query("id")
	field := c.Query("field")
	fieldContent := c.Query("fieldContent")

	err := new(models.Job).UpdateFieldContent(id, field, fieldContent)

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
		"msg":    "update completed",
	})
}

/**
流程
任务是否加锁->仓库状态是否正常->仓库是否忙碌->
任务状态进行中->更新仓库状态为繁忙->响应前端->
更新代码->检查分支->切换分支->更新代码->安装依赖->
打包->创建目录并复制代码->更新Url->
更新运行次数->更新任务状态->更新仓库状态为空闲->（执行附加脚本，暂不做）->更新终端信息。
*/
func JobRun(c *gin.Context) {

	jobId := c.Query("id")

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
	if job.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
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
			"msg":    "The repository is damaged or is being cloned, please try to repair the corresponding repository",
		})
		return
	}

	if repository.JobStatus != models.RepoJobStatusLeisured {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Repository resources have been occupied, please wait a while and try again",
		})
		return
	}
	//
	claims, _ := serviceUser.GetJwtClaimsByContext(c)
	msg := models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” run task“" + job.Name + "”",
	}
	models.Broadcast <- msg

	//开始执行任务，任务状态和仓库的任务状态
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

	go serviceJob.JobRun(&job, &repository, claims)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Compiling in the background",
		"msg":    "",
	})

}

func JobLock(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	claims, _ := serviceUser.GetJwtClaimsByContext(c)

	var job models.Job
	if sql.DB.First(&job, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	//加锁
	if job.LockPassword == "" {
		if sql.DB.Model(&job).
			UpdateColumn("lock_password", password).
			UpdateColumn("user", claims.Username).
			Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    sql.DB.Error.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"data":   "",
				"msg":    "Task locked successfully",
			})
			return
		}

	}
	//解锁
	if job.LockPassword != password {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Incorrect unlock password",
		})
		return
	}
	if sql.DB.Model(&job).UpdateColumn("lock_password", "").Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "Mission unlocked successfully",
	})

}
