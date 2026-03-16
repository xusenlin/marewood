package context

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	AppErr "marewood/internal/pkg/errors"
)

func (r *Context) SendErr(err error) {
	status := http.StatusInternalServerError

	var appErr *AppErr.AppError
	if errors.As(err, &appErr) {
		status = appErr.HttpStatus
	}

	r.c.JSON(status, gin.H{
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
