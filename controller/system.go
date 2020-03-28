package controller

import (
	"MareWood/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": map[string]interface{}{
			"kvConfig":        config.Cfg,
			"detailsConfig": config.GetSysInfo(),
		},
		"msg": "查询成功",
	})
}
