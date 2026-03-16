package routes

import (
	"github.com/gin-gonic/gin"
	"marewood/conf"
	"marewood/internal/controllers"
	"marewood/internal/dao"
	"marewood/internal/pkg/context"
	"marewood/internal/pkg/database"
	"marewood/internal/services"
	"marewood/middlewares"
)

func setV1Api(r *gin.Engine) {
	// 获取数据库连接
	db := database.GetDB()

	// 初始化 DAOs
	userDAO := dao.NewUserDAO(db)
	repoDAO := dao.NewRepositoryDAO(db)
	taskDAO := dao.NewTaskDAO(db)

	// 初始化 Services
	gitService := services.NewGitService()
	userService := services.NewUserService(userDAO)
	repoService := services.NewRepositoryService(repoDAO, gitService)
	taskService := services.NewTaskService(taskDAO, repoDAO, gitService)

	// 初始化 Controllers
	userCtrl := controllers.NewUserController(userService)
	repoCtrl := controllers.NewRepositoryController(repoService)
	taskCtrl := controllers.NewTaskController(taskService)

	// 公开路由
	v1Public := r.Group("/v1").Use(middlewares.Logger())
	{
		v1Public.POST("/login", userCtrl.Login)
		v1Public.POST("/register", userCtrl.Register)

		v1Public.GET("/event/repository", repoCtrl.EventSource)
		v1Public.GET("/event/task", taskCtrl.EventSource)
		v1Public.GET("/task/archiver", taskCtrl.Archiver)
	}

	// 需要认证的路由
	v1 := r.Group("/v1").Use(middlewares.JWTAuth(), middlewares.AuthLogger())
	{
		v1.GET("/sys", func(c *gin.Context) {
			ctx := context.New(c)
			ctx.SendOk("", conf.ConfigMap())
		})

		// 用户路由
		v1.GET("/users", middlewares.RoleReporter(), userCtrl.Find)
		v1.POST("/user", middlewares.RoleSuperAdmin(), userCtrl.Edit)
		v1.DELETE("/user/:id", middlewares.RoleSuperAdmin(), userCtrl.Destroy)

		// 仓库路由
		v1.GET("/repositories", middlewares.RoleReporter(), repoCtrl.Find)
		v1.GET("/repository/list", middlewares.RoleReporter(), repoCtrl.FindAll)
		v1.PUT("/repository", middlewares.RoleDeveloper(), repoCtrl.Edit)
		v1.DELETE("/repository/:id", middlewares.RoleSuperAdmin(), repoCtrl.Destroy)
		v1.GET("/repository/git_pull/:id", middlewares.RoleDeveloper(), repoCtrl.GitPull)
		v1.GET("/repository/git_checkout_dot/:id", middlewares.RoleDeveloper(), repoCtrl.GitCheckoutDot)
		v1.GET("/repository/branch/:repositoryId", middlewares.RoleReporter(), repoCtrl.Branch)
		v1.GET("/repository/package_json/:repositoryId", middlewares.RoleReporter(), repoCtrl.GetPackageJson)

		// 任务路由
		v1.GET("/tasks", middlewares.RoleReporter(), taskCtrl.Find)
		v1.GET("/task/tags", middlewares.RoleReporter(), taskCtrl.Tags)
		v1.PUT("/task", middlewares.RoleDeveloper(), taskCtrl.Edit)
		v1.GET("/task/update_branch", middlewares.RoleDeveloper(), taskCtrl.UpdateBranch)
		v1.DELETE("/task/:id", middlewares.RoleDeveloper(), taskCtrl.Destroy)
		v1.GET("/task/run", middlewares.RoleDeveloper(), taskCtrl.Run)
	}
}
