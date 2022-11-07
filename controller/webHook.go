package controller

import (
	"MareWood/config"
	"MareWood/helper"
	"MareWood/models"
	"MareWood/service/serviceTask"
	"MareWood/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

//webHook 触发任务
func TaskWebHook(c *gin.Context) {
	TaskRun(c)
}

func RunTaskAndPack(c *gin.Context) {
	taskId := c.Query("id")

	var task models.Task
	var repository models.Repository

	if taskId == "" {
		c.Data(http.StatusBadRequest, "text", []byte("taskId is empty"))
		return
	}

	if sql.DB.First(&task, taskId).Error == gorm.ErrRecordNotFound {
		c.Data(http.StatusBadRequest, "text", []byte(gorm.ErrRecordNotFound.Error()))
		return
	}
	if sql.DB.Error != nil {
		c.Data(http.StatusBadRequest, "text", []byte(sql.DB.Error.Error()))
		return
	}

	if task.LockPassword != "" {
		c.Data(http.StatusBadRequest, "text", []byte(sql.DB.Error.Error()))
		return
	}
	if sql.DB.First(&repository, task.RepositoryId).Error != nil {
		c.Data(http.StatusBadRequest, "text", []byte(sql.DB.Error.Error()))
		return
	}
	if repository.Status != models.RepoStatusSuccess {
		c.Data(http.StatusBadRequest, "text", []byte("The repository is damaged or is being cloned, please try to repair the corresponding repository"))
		return
	}

	if repository.TaskStatus != models.RepoTaskStatusLeisured {
		c.Data(http.StatusBadRequest, "text", []byte("Repository resources have been occupied, please wait a while and try again"))
		return
	}
	//开始执行任务，任务状态和仓库的任务状态
	repository.TaskStatus = models.RepoTaskStatusBusy
	task.Status = models.TaskStatusProcessing

	if sql.DB.Save(&repository).Save(&task).Error != nil {
		c.Data(http.StatusBadRequest, "text", []byte(sql.DB.Error.Error()))
		return
	}

	//go
	var emptyClaims *models.Claims

	serviceTask.TaskRun(&task, &repository, emptyClaims)

	if sql.DB.First(&task, taskId).Error != nil {
		c.Data(http.StatusBadRequest, "text", []byte(sql.DB.Error.Error()))
		return
	}

	if task.Status != models.TaskStatusSuccess {
		c.Data(http.StatusBadRequest, "text", []byte(task.TerminalInfo))
		return
	}

	var src = config.Cfg.WebRootDir + "/" + taskId
	var dst = fmt.Sprintf("%s.tar.gz", src)

	// 将步骤写入了一个函数中，这样处理错误方便一些
	if err := helper.Tar(src, dst); err != nil {
		c.Data(http.StatusBadRequest, "text", []byte(err.Error()))
		return
	}

	c.File(dst)

	return
}
