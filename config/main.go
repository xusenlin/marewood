package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

	b, err := ioutil.ReadFile(systemCfg.CurrentDir + "/config.json")

	if err != nil {
		panic(err)
	}

	if json.Unmarshal(b, &userCfg) != nil {
		panic(err)
	}

	ClientDir := os.Getenv("ClientDir")
	if ClientDir == "" {
		ClientDir = systemCfg.CurrentDir + "/public/build"
	}

	systemCfg.Version = 0.4
	systemCfg.AppName = "MareWood"
	systemCfg.AppRepository = "https://github.com/xusenlin/MareWood"
	systemCfg.DbDns = systemCfg.CurrentDir + "/database.db"
	systemCfg.ResourcesDir = systemCfg.CurrentDir + "/resources"
	systemCfg.RepositoryDir = systemCfg.ResourcesDir + "/repositories"
	systemCfg.WebRootDir = systemCfg.ResourcesDir + "/webs"
	systemCfg.ClientDir = ClientDir
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
