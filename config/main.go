package config

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type Config struct {
	userConfig
	systemConfig
}

var systemCfg systemConfig
var userCfg userConfig
var Cfg Config

func init() {
	var err error
	if systemCfg.CurrentDir, err = os.Getwd(); err != nil {
		panic(err)
	}

	userCfg.DependTools = []string{"npm", "yarn"}
	userCfg.LogMode = false

	userCfg.GinMode = os.Getenv("MW_GIN_MODE")
	if userCfg.GinMode == "" {
		userCfg.GinMode = gin.ReleaseMode
	}

	userCfg.HttpPort = os.Getenv("MW_PORT")
	if userCfg.HttpPort == "" {
		userCfg.HttpPort = "8088"
	}

	otherAddressUrl := os.Getenv("MW_URL")
	if otherAddressUrl != "" {
		userCfg.OtherAddressUrl = strings.Split(otherAddressUrl, ",")
	}

	systemCfg.Version = 0.40
	systemCfg.AppName = "MareWood"
	systemCfg.AppRepository = "https://github.com/xusenlin/MareWood"
	systemCfg.ResourcesDir = systemCfg.CurrentDir + "/resources"
	systemCfg.DbDns = systemCfg.ResourcesDir + "/database.db"
	systemCfg.RepositoryDir = systemCfg.ResourcesDir + "/repositories"
	systemCfg.WebRootDir = systemCfg.ResourcesDir + "/webs"
	systemCfg.ClientDir = systemCfg.CurrentDir + "/public/build"
	systemCfg.WebHookUrl = "/web_hook"
	systemCfg.WebsUrl = "/webs"
	Cfg.systemConfig = systemCfg
	Cfg.userConfig = userCfg
}

func GetSysInfo() []interface{} {
	type rowInfo struct {
		Title string
		Val   interface{}
	}
	return []interface{}{
		rowInfo{"Name", Cfg.AppName},
		rowInfo{"Version", Cfg.Version},
		rowInfo{"RunDir", Cfg.CurrentDir},
		rowInfo{"DatabaseAddress", Cfg.DbDns},
		rowInfo{"ResourcesDir", Cfg.ResourcesDir},
		rowInfo{"RepositoryDir", Cfg.RepositoryDir},
		rowInfo{"WebRootDir", Cfg.WebRootDir},
		rowInfo{"ClientDir", Cfg.ClientDir},
		rowInfo{"DependTools", Cfg.DependTools},
		rowInfo{"WebHookUrl", Cfg.WebHookUrl},
		rowInfo{"WebsUrl", Cfg.WebsUrl},
		rowInfo{"GitHub", Cfg.AppRepository},
	}
}
