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
				"msg":    "Only reporters and above roles have permission to operate",
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
				"msg":    "Developer and above roles have the right to operate",
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
				"msg":    "Only the administrator has the authority to operate",
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
				"msg":    "Only super administrators have authority to operate",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
