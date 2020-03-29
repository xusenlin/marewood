package middlewares

import (
	"MareWood/models"
	"MareWood/service/serviceUser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleReporter() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := serviceUser.GetJwtClaimsByContext(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Role > models.UserRoleReporter {
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
		claims, err := serviceUser.GetJwtClaimsByContext(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Role > models.UserRoleDeveloper {
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
		claims, err := serviceUser.GetJwtClaimsByContext(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Role > models.UserRoleAdminister {
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
		claims, err := serviceUser.GetJwtClaimsByContext(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Role > models.UserRoleSuperAdministrator {
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
