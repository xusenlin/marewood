package routes

import (
	"FrontEndAutomatedDeployment/config"
	"FrontEndAutomatedDeployment/controller"
	"FrontEndAutomatedDeployment/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.Static("/public", config.Cfg.ClientDir)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/public")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/categories", controller.CategoryFindAll)
		v1.POST("/category/create", controller.CategoryCreate)
		v1.GET("/category/delete", controller.CategoryDestroy)
		//v1.POST("/user/update", Controllers.UserUpdate)
		//v1.POST("/users", Controllers.UserFindAll)
	}

	return r
}
