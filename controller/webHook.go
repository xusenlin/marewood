package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//webHook 触发任务
func JobWebHook(c *gin.Context) {

	repositoryId := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   repositoryId,
		"msg":    "触发任务还没有完成",
	})
	return
}
