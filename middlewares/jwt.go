package middlewares

import (
	"marewood/internal/pkg/context"
	AppErr "marewood/internal/pkg/errors"
	"marewood/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		ctx := context.New(c)
		if token == "" {
			ctx.SendErr(AppErr.NewForbiddenError("please log in first"))
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
