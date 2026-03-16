package context

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	c *gin.Context
}

func New(c *gin.Context) *Context {
	return &Context{c: c}
}
