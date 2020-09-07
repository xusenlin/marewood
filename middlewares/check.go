package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAfferent(keyList []string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			lens               = len(keyList)
			isNext      bool   = false
			unResParams string = ""
		)
		if lens <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "参数错误",
			})
			c.Abort()
			return
		}
		for _, item := range keyList {
			var value string = c.Query(item)
			if value == "" {
				if unResParams == "" {
					unResParams += item
				} else {
					unResParams += "," + item
				}
				isNext = true
			}
		}
		if !isNext {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "缺失" + unResParams,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
