package context

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	AppErr "marewood/internal/pkg/errors"
)

func (r *Context) SendErr(err error) {
	errCode := AppErr.ErrCodeUndefined

	var appErr *AppErr.AppError
	if errors.As(err, &appErr) {
		errCode = appErr.Code
	}

	r.c.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   nil,
		"code":   errCode,
		"msg":    err.Error(),
	})
}

func (r *Context) SendOk(msg string, data any) {
	r.c.JSON(http.StatusOK, gin.H{
		"status": true,
		"code":   200,
		"data":   data,
		"msg":    msg,
	})
}
