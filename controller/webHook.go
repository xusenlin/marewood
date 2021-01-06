package controller

import (
	"github.com/gin-gonic/gin"
)

//webHook 触发任务
func TaskWebHook(c *gin.Context) {
	TaskRun(c)
}
