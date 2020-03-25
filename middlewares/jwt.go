package middlewares

import (
	"MareWood/helper"
	"MareWood/models"
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
		claims, err := helper.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(models.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}