package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"marewood/internal/context"
	"marewood/internal/user"
)

func RoleReporter() func(c *gin.Context) {
	tip := errors.New("only reporters and above roles have permission to operate")
	return func(c *gin.Context) { role(c, user.RoleReporter, tip) }
}

func RoleDeveloper() func(c *gin.Context) {
	tip := errors.New("developer and above roles have the right to operate")
	return func(c *gin.Context) { role(c, user.RoleDeveloper, tip) }
}

func RoleAdmin() func(c *gin.Context) {
	tip := errors.New("only the administrator has the authority to operate")
	return func(c *gin.Context) { role(c, user.RoleAdminister, tip) }
}

func RoleSuperAdmin() func(c *gin.Context) {
	tip := errors.New("only super administrators have authority to operate")
	return func(c *gin.Context) { role(c, user.RoleSuperAdministrator, tip) }
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
