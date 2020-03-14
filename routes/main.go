package routes

import (
	"FEDeployService/config"
	"FEDeployService/controller"
	"FEDeployService/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	r.Use(middlewares.Cors())

	{
		r.GET("/ping", func (c *gin.Context) {//服务健康检查
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		r.Static(config.Cfg.WebsUrl, config.Cfg.WebRootDir)
		r.Static("/public", config.Cfg.ClientDir)

		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/public")
		})

		r.POST(config.Cfg.WebHookUrl, controller.JobWebHook)
	}


	v1 := r.Group("/v1")
	{
		//系统信息
		v1.GET("/system/info", controller.SystemInfo)
		//仓库相关
		v1.GET("/repositories", controller.RepositoryFindAll)
		v1.POST("/repository/create", controller.RepositoryCreate)
		v1.GET("/repository/delete", controller.RepositoryDestroy)
		v1.GET("/repository/git_pull", controller.RepositoryGitPull)
		v1.GET("/repository/git_branch", controller.RepositoryBranch)
		v1.GET("/repository/delete_depend", controller.RepositoryDeleteDepend)
		//任务分类
		v1.GET("/categories", controller.CategoryFindAll)
		v1.POST("/category/create", controller.CategoryCreate)
		v1.GET("/category/delete", controller.CategoryDestroy)
		//任务
		v1.GET("/jobs", controller.JobFindAll)
		v1.GET("/jobs_find", controller.JobFindByCategoryId)
		v1.POST("/job/create", controller.JobCreate)
		v1.GET("/job/update_branch", controller.JobUpdateBranch)
		v1.GET("/job/delete", controller.JobDestroy)
		v1.GET("/job/run", controller.JobRun)

	}

	return r
}
