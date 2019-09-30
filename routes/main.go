package routes

import (
	"FrontEndAutomatedDeployment/config"
	"github.com/gin-gonic/gin"
)

//func InitRoute()  {
//	initWebRoute()
//	initApiRoute()
//}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/", config.Cfg.ClientDir) // 静态文件服务
	//r.LoadHTMLGlob("views/**/*") // 载入html模板目录

	// web路由

	//r.GET("/about", Controllers.About)
	//r.GET("/post/index", Controllers.Post)
	//
	//// 简单的路由组: v1
	//v1 := r.Group("/api")
	//{
	//	v1.GET("/ping", Controllers.Ping)
	//	v1.POST("/user/create", Controllers.UserCreate)
	//	v1.POST("/user/delete", Controllers.UserDestroy)
	//	v1.POST("/user/update", Controllers.UserUpdate)
	//	v1.POST("/users", Controllers.UserFindAll)
	//}

	return r
}
