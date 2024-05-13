package task

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
	"gorm.io/gorm"
	"io"
	"log/slog"
	"marewood/internal/common"
	"marewood/internal/context"
	"marewood/internal/db"
	"marewood/internal/event"
	"marewood/internal/log"
	"strconv"
	"strings"
)

func Find(c *gin.Context) {

	ctx := context.New(c)
	model := new(Task)

	claims, err := ctx.GetClaims() //claims.ID

	if err != nil {
		ctx.SendErr(err)
		return
	}

	pagination := ginPagination.New[*Task](model, c)
	pagination.Eq("id").Like("name")

	tagsStr := strings.Trim(c.Query("tags"), " ")
	if tagsStr != "" {
		tags := strings.Split(tagsStr, ",")
		pagination.CB(func(db *gorm.DB) {
			db.Where("tag in ?", tags)
		})
	}
	pagination.CB(func(db *gorm.DB) {
		db.Where("(private = ? AND created_id = ?) OR private = ?", true, claims.ID, false)
	})

	err = pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("pagination success", pagination)
}

func Tags(c *gin.Context) {
	var tags []struct {
		Tag   string `json:"tag"`
		Count int    `json:"count"`
	}
	ctx := context.New(c)
	err := db.Conn.Model(Task{}).
		Select("tag, count(*) as count").
		Group("tag").
		Find(&tags).Error
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("success", tags)
}

func Edit(c *gin.Context) {

	var task Task
	ctx := context.New(c)
	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = c.ShouldBindJSON(&task)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	err = task.Edit(claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	event.TaskSource.PublishMsgExcludeID(
		event.TaskTypeEditOk,
		&event.TaskData{
			TaskId:   task.ID,
			UserId:   claims.ID,
			UserName: claims.Username,
			Msg:      "edit success",
		},
		claims.ID,
	)
	ctx.SendOk("edit success", task.ID)
}

func UpdateBranch(c *gin.Context) {
	ctx := context.New(c)

	id := c.Query("id")
	branch := c.Query("branch")

	if id == "" || branch == "" {
		ctx.SendErr(errors.New("id or branch is not allowed to be empty"))
		return
	}

	err := db.Conn.Model(Task{}).Where("id = ?", id).Update("branch", branch).Error
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("edit success", id)
}

func EventSource(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	claims, err := common.JwtParseToken(c.Query("token"))

	if err != nil {
		c.Stream(func(w io.Writer) bool {
			c.SSEvent("error", "token error")
			return false
		})
		return
	}

	event.TaskSource.Subscribe(claims.ID)
	defer event.TaskSource.CancelSubscribe(claims.ID)
	log.Slog.Info("[TaskEventSource]:subscribe", slog.Any("userId", claims.ID))
	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-event.TaskSource.ReceiveMsg(claims.ID):
			c.SSEvent(msg.Type, msg.Data)
			log.Slog.Info("[TaskEventSource]:sendMsg", slog.Any("userId", claims.ID), slog.String("type", msg.Type), slog.String("msg", msg.Data.Msg))
			return true
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}

func Archiver(c *gin.Context) {
	ctx := context.New(c)
	id := c.Query("id")
	archiverType := c.Query("type")
	if id == "" {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}
	var task Task
	if err := db.Conn.First(&task, id).Error; err != nil {
		ctx.SendErr(err)
		return
	}

	if task.Status != StatusSuccess {
		ctx.SendErr(errors.New("task status is not successful"))
		return
	}

	switch archiverType {
	case "1":
		file, err := task.Tar()
		if err != nil {
			ctx.SendErr(err)
			return
		}
		ctx.SendFile(fmt.Sprintf("%s.tar", task.CommitHash), file)
		return
	case "2":
		file, err := task.Zip()
		if err != nil {
			ctx.SendErr(err)
			return
		}
		ctx.SendFile(fmt.Sprintf("%s.zip", task.CommitHash), file)
		return
	default:
		ctx.SendErr(errors.New("type is not allowed to be empty"))
		return
	}

}

func Destroy(c *gin.Context) {
	ctx := context.New(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}

	var task Task
	task.ID = uint(id)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = task.Destroy(claims)

	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("successfully deleted", id)
}
