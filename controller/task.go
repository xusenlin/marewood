package controller

import (
	"MareWood/models"
	"MareWood/service/serviceTask"
	"MareWood/service/serviceUser"
	"MareWood/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func TaskFindAll(c *gin.Context) {

	result, err := new(models.Task).FindAll()

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

func TaskFind(c *gin.Context) {

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
	var result models.TaskPageResult
	if strings.HasPrefix(name, "id:") {
		result, err = models.FindTask("", pageNum, pageSize,
			map[string]interface{}{"id": strings.TrimLeft(name, "id:"), "category_id": categoryId})
	} else {
		result, err = models.FindTask(name, pageNum, pageSize, map[string]interface{}{"category_id": categoryId})
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

func TaskCreate(c *gin.Context) {

	var task models.Task
	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err = task.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err = new(models.Category).CategoryTaskQuantityIncrement(task.CategoryId)
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

func TaskUpdateBranch(c *gin.Context) {

	id := c.Query("id")
	branch := c.Query("branch")

	var task models.Task

	if sql.DB.First(&task, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if task.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
		})
		return
	}
	if err := task.UpdateBranch(branch); err != nil {
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

func TaskDestroy(c *gin.Context) {

	id := c.Query("id")
	var task models.Task

	if sql.DB.First(&task, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if task.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
		})
		return
	}

	err := task.Destroy(id)
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
		Message:         "“" + claims.Username + "”deleted the task “" + task.Name + "”",
	}
	models.Broadcast <- msg
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
	})
}

func TaskUpdateField(c *gin.Context) {

	id := c.Query("id")
	field := c.Query("field")
	fieldContent := c.Query("fieldContent")

	err := new(models.Task).UpdateFieldContent(id, field, fieldContent)

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
func TaskRun(c *gin.Context) {

	taskId := c.Query("id")

	if taskId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "taskId is empty",
		})
		return
	}

	var task models.Task
	var repository models.Repository

	if sql.DB.First(&task, taskId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if task.LockPassword != "" {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Please unlock the task first",
		})
		return
	}
	if sql.DB.First(&repository, task.RepositoryId).Error != nil {
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

	if repository.TaskStatus != models.RepoTaskStatusLeisured {
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
		Message:         "“" + claims.Username + "” run task“" + task.Name + "”",
	}
	models.Broadcast <- msg

	//开始执行任务，任务状态和仓库的任务状态
	repository.TaskStatus = models.RepoTaskStatusBusy
	task.Status = models.TaskStatusProcessing

	if sql.DB.Save(&repository).Save(&task).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	//go

	go serviceTask.TaskRun(&task, &repository, claims)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Compiling in the background",
		"msg":    "",
	})

}

func TaskLock(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	claims, _ := serviceUser.GetJwtClaimsByContext(c)

	var task models.Task
	if sql.DB.First(&task, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	//加锁
	if task.LockPassword == "" {
		if sql.DB.Model(&task).
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
	if task.LockPassword != password {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Incorrect unlock password",
		})
		return
	}
	if sql.DB.Model(&task).UpdateColumn("lock_password", "").Error != nil {
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
