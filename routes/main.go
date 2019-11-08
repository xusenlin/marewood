package routes

import (
	"FEDeployService/config"
	"FEDeployService/controller"
	"FEDeployService/middlewares"
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

	r.Any(config.Cfg.WebHookUrl, controller.RepositoryUpdate) //WebHook仓库更新

	v1 := r.Group("/v1")
	{
		v1.GET("/categories", controller.CategoryFindAll)
		v1.POST("/category/create", controller.CategoryCreate)
		v1.GET("/category/delete", controller.CategoryDestroy)
		//v1.POST("/user/update", Controllers.UserUpdate)

		v1.GET("/repositories", controller.RepositoryFindAll)
		v1.POST("/repository/create", controller.RepositoryCreate)
		v1.GET("/repository/delete", controller.RepositoryDestroy)
		//v1.POST("/user/update", Controllers.UserUpdate)


	}

	return r
}
