package context

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Context) SendErr(err error) {
	r.c.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   "",
		"msg":    err.Error(),
	})
}

func (r *Context) SendOk(msg string, data any) {
	r.c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   data,
		"msg":    msg,
	})
}

