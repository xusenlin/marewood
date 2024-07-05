package routes

import (
	"github.com/gin-gonic/gin"
	"marewood/conf"
	"marewood/internal/context"
	"marewood/internal/repository"
	"marewood/internal/task"
	"marewood/internal/user"
	"marewood/middlewares"
)

func setV1Api(r *gin.Engine) {
	v1Public := r.Group("/v1").Use(middlewares.Logger())
	{
		v1Public.POST("/login", user.Login)
		v1Public.POST("/register", user.Register)

		v1Public.GET("/event/repository", repository.EventSource)
		v1Public.GET("/event/task", task.EventSource)
		v1Public.GET("/task/archiver", task.Archiver)
	}

	v1 := r.Group("/v1").Use(middlewares.JWTAuth(), middlewares.AuthLogger())
	{
		v1.GET("/sys", func(c *gin.Context) {
			ctx := context.New(c)
			ctx.SendOk("", conf.ConfigMap())
		})
		//user
		v1.GET("/users", middlewares.RoleReporter(), user.Find)
		v1.POST("/user", middlewares.RoleSuperAdmin(), user.Edit)
		v1.DELETE("/user/:id", middlewares.RoleSuperAdmin(), user.Destroy)
		//repository
		v1.GET("/repositories", middlewares.RoleReporter(), repository.Find)
		v1.GET("/repository/list", middlewares.RoleReporter(), repository.FindAll)
		v1.PUT("/repository", middlewares.RoleDeveloper(), repository.Edit)
		v1.DELETE("/repository/:id", middlewares.RoleSuperAdmin(), repository.Destroy)
		v1.GET("/repository/git_pull/:id", middlewares.RoleDeveloper(), repository.GitPull)
		v1.GET("/repository/git_checkout_dot/:id", middlewares.RoleDeveloper(), repository.GitCheckoutDot)
		v1.GET("/repository/branch/:repositoryId", middlewares.RoleReporter(), repository.Branch)
		v1.GET("/repository/package_json/:repositoryId", middlewares.RoleReporter(), repository.GetPackageJson)

		//task
		v1.GET("/tasks", middlewares.RoleReporter(), task.Find)
		v1.GET("/task/tags", middlewares.RoleReporter(), task.Tags)
		v1.PUT("/task", middlewares.RoleDeveloper(), task.Edit)
		v1.GET("/task/update_branch", middlewares.RoleDeveloper(), task.UpdateBranch)
		v1.DELETE("/task/:id", middlewares.RoleDeveloper(), task.Destroy)
		v1.GET("/task/run", middlewares.RoleDeveloper(), task.Run)
	}
}
