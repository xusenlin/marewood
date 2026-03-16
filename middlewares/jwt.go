package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"marewood/internal/pkg/context"
	"marewood/internal/pkg/jwt"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		ctx := context.New(c)
		if token == "" {
			ctx.SendErr(errors.New("please log in first"))
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			ctx.SendErr(err)
			c.Abort()
			return
		}
		c.Set(jwt.JwtClaimsKey, claims)
		c.Next() // 后续的处理函数可以用过c.Get("JwtClaims")来获取当前请求的用户信息
	}
}
