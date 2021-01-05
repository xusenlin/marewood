package controller

import (
	"github.com/gin-gonic/gin"
)

//webHook 触发任务
func JobWebHook(c *gin.Context) {
	JobRun(c)
}
