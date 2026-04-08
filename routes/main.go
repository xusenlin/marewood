package routes

import (
	"marewood/conf"
	"marewood/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(conf.GinMode)
	r := gin.Default()

	r.Use(middlewares.Cors())

	{
		r.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		r.Static("/webs", conf.WebRootDir)
		r.Static("/public", conf.ClientDir)

		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/public")
		})
	}

	setV1Api(r)

	return r
}
