package middlewares

import (
	"MareWood/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleReporter() func(c *gin.Context) {
	return func(c *gin.Context) {
		role := c.GetInt("JWT_ROLE")
		if role > models.UserRoleReporter {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "记者及以上角色才有权限操作",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RoleDeveloper() func(c *gin.Context) {
	return func(c *gin.Context) {
		role := c.GetInt("JWT_ROLE")
		if role > models.UserRoleDeveloper {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "开发者及以上角色才有权限操作",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}


func RoleAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		role := c.GetInt("JWT_ROLE")
		if role > models.UserRoleAdminister {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "管理员才有权限操作",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RoleSuperAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		role := c.GetInt("JWT_ROLE")
		if role > models.UserRoleSuperAdministrator {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "超级管理员才有权限操作",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}