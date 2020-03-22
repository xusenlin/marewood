package middlewares

import (
	"MareWood/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)


func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "请先登录",
			})
			c.Abort()
			return
		}
		userInfo, err := helper.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("JWT_ID", userInfo.ID)
		c.Set("JWT_NAME", userInfo.Username)
		c.Set("JWT_ROLE", userInfo.Role)
		c.Set("JWT_STATUS", userInfo.Status)
		c.Next() // 后续的处理函数可以用过c.Get("JWT_ROLE")来获取当前请求的用户信息
	}
}