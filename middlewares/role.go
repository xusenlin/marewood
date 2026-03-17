package middlewares

import (
	"marewood/internal/models"
	"marewood/internal/pkg/context"
	AppErr "marewood/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

func RoleReporter() func(c *gin.Context) {
	tip := AppErr.NewUnauthorizedError("only reporters and above roles have permission to operate")
	return func(c *gin.Context) { role(c, models.UserRoleReporter, tip) }
}

func RoleDeveloper() func(c *gin.Context) {
	tip := AppErr.NewUnauthorizedError("developer and above roles have the right to operate")
	return func(c *gin.Context) { role(c, models.UserRoleDeveloper, tip) }
}

func RoleSuperAdmin() func(c *gin.Context) {
	tip := AppErr.NewUnauthorizedError("only super administrators have authority to operate")
	return func(c *gin.Context) { role(c, models.UserRoleSuperAdmin, tip) }
}

func role(c *gin.Context, roleVal int, errTip error) {
	ctx := context.New(c)
	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		c.Abort()
		return
	}
	if claims.Role > roleVal {
		ctx.SendErr(errTip)
		c.Abort()
		return
	}
	c.Next()
}
