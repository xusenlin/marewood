package controllers

import (
	"io"
	"marewood/internal/models"
	"marewood/internal/pkg/context"
	"marewood/internal/pkg/errors"
	"marewood/internal/pkg/event"
	"marewood/internal/pkg/jwt"
	"marewood/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
)

// TaskController 任务控制器
type TaskController struct {
	taskService services.TaskService
}

// NewTaskController 创建任务控制器
func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

// Find 查询任务列表（分页）
func (ctrl *TaskController) Find(c *gin.Context) {
	ctx := context.New(c)

	model := new(models.Task)

	_, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	pagination := ginPagination.New[*models.Task](model, c)
	pagination.Eq("id").Like("name")

	err = pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("pagination success", pagination)
}

// Tags 获取所有标签
func (ctrl *TaskController) Tags(c *gin.Context) {
	ctx := context.New(c)

	tags, err := ctrl.taskService.GetAllTags()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("get tags success", tags)
}

// Edit 创建或编辑任务
func (ctrl *TaskController) Edit(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	var task models.Task
	err = c.ShouldBindJSON(&task)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	err = ctrl.taskService.CreateOrUpdate(&task, claims.ID)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("edit success", task.ID)
}

// UpdateBranch 更新任务分支
func (ctrl *TaskController) UpdateBranch(c *gin.Context) {
	ctx := context.New(c)

	taskId, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil || taskId == 0 {
		ctx.SendErr(errors.NewValidationError("id is not allowed to be empty"))
		return
	}

	repositoryId, err := strconv.ParseUint(c.Query("repositoryId"), 10, 32)
	if err != nil || repositoryId == 0 {
		ctx.SendErr(errors.NewValidationError("repositoryId is not allowed to be empty"))
		return
	}

	err = ctrl.taskService.UpdateBranch(uint(taskId), uint(repositoryId))
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("update branch success", taskId)
}

// Destroy 删除任务
func (ctrl *TaskController) Destroy(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.SendErr(errors.NewValidationError("id is not allowed to be empty"))
		return
	}

	err = ctrl.taskService.DeleteTask(uint(id), claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("successfully deleted", id)
}

// Run 运行任务
func (ctrl *TaskController) Run(c *gin.Context) {
	ctx := context.New(c)

	taskId := c.Query("id")
	if taskId == "" {
		ctx.SendErr(errors.NewValidationError("id is not allowed to be empty"))
		return
	}

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	id, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		ctx.SendErr(errors.NewValidationError("invalid id"))
		return
	}

	err = ctrl.taskService.RunTask(uint(id), claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("Compiling in the background...", id)
}

// Archiver 下载任务打包文件
func (ctrl *TaskController) Archiver(c *gin.Context) {
	ctx := context.New(c)

	taskId := c.Query("id")
	if taskId == "" {
		ctx.SendErr(errors.NewValidationError("id is not allowed to be empty"))
		return
	}
	format := c.Query("format")
	id, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		ctx.SendErr(errors.NewValidationError("invalid id"))
		return
	}

	data, fileName, err := ctrl.taskService.ArchiveTask(uint(id), format)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	// 返回文件
	switch format {
	case "zip":
		ctx.SendZipFile(fileName, data)
	case "tar":
		ctx.SendTarFile(fileName, data)
	default:
		ctx.SendErr(errors.NewValidationError("unsupported archive format"))
	}

}

// EventSource SSE 事件流
func (ctrl *TaskController) EventSource(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	claims, err := jwt.ParseToken(c.Query("token"))
	var userID uint
	if err != nil {
		userID = 0
	} else {
		userID = claims.ID
	}

	event.TaskSource.Subscribe(userID)
	c.SSEvent("message", "connected")
	c.Writer.Flush()
	defer event.TaskSource.CancelSubscribe(userID)

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-event.TaskSource.ReceiveMsg(userID):
			c.SSEvent(msg.Type, msg.Data)
			return true
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}
