package controller

import (
	"FEDeployService/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DependentSupport(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   config.Cfg.DependentTools,
		"msg":    "成功",
	})
}

