package middlewares

import (
	"FEDeployService/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)


func JWTAuthMiddleware() func(c *gin.Context) {
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
				"msg":    "无效的 Token",
			})
			c.Abort()
			return
		}
		c.Set("user", userInfo)
		c.Next() // 后续的处理函数可以用过c.Get("user")来获取当前请求的用户信息
	}
}