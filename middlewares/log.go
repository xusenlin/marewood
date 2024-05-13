package middlewares

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"marewood/internal/common"
	"marewood/internal/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		statusCode := c.Writer.Status()

		log.GIN.Info(
			"[GIN] request middlewares",
			slog.String("clientIP", clientIP),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("statusCode", statusCode),
		)
		c.Next()
	}
}

func AuthLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		statusCode := c.Writer.Status()

		user := new(common.Claims)

		jwtClaims, hasClaims := c.Get(common.JwtClaimsKey)
		if hasClaims {
			claims, ok := jwtClaims.(*common.Claims)
			if ok {
				user = claims
			}
		}

		log.GIN.Info(
			"[GIN] request middlewares",
			slog.String("clientIP", clientIP),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("statusCode", statusCode),
			slog.Any("userID", user.ID),
			slog.String("userName", user.Username),
			slog.Int("userRole", user.Role),
			slog.Int("userStatus", user.Status),
		)
		c.Next()
	}
}
