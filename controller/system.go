package controller

import (
	"FEDeployService/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemInfo(c *gin.Context) {
	kvConfig, detailsConfig := config.GetSysInfo();
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": map[string]interface{}{
			"kvConfig":        kvConfig,
			"detailsConfig": detailsConfig,
		},
		"msg": "查询成功",
	})
}
