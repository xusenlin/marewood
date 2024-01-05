package conf

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	HttpPort string
	GinMode  string
)

func init() {
	GinMode = os.Getenv("MW_GIN_MODE")
	if GinMode == "debug" {
		GinMode = gin.DebugMode
	} else {
		GinMode = gin.ReleaseMode
	}

	HttpPort = os.Getenv("MW_PORT")
	if HttpPort == "" {
		HttpPort = "8088"
	}
}
