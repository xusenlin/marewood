package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type userConfig struct {
	Port int `json:"port"`
}

type Config struct {
	userConfig
	DbDns          string
	Version        float32
	AppName        string
	CurrentDir     string
	AppRepository  string
	ConfigFilePath string
	RepositoryDir  string
	WebRootDir     string
	ClientDir      string
}

var Cfg Config

func init() {
	var err error
	if Cfg.CurrentDir, err = os.Getwd(); err != nil {
		panic(err)
	}

	Cfg.Port = 8080
	Cfg.Version = 0.1
	Cfg.AppName = "前端自动部署工具"
	Cfg.AppRepository = ""
	Cfg.ConfigFilePath = Cfg.CurrentDir + "/config.json"
	Cfg.DbDns = Cfg.CurrentDir + "/database.db"
	Cfg.RepositoryDir = Cfg.CurrentDir + "/repository" //克隆过来的所有仓库
	Cfg.WebRootDir = Cfg.CurrentDir + "/webs"          //打包完成后的静态文件存放地址
	Cfg.ClientDir = Cfg.CurrentDir + "/public/build"   //客户端地址

	var userCgf userConfig

	configFile, err := ioutil.ReadFile(Cfg.ConfigFilePath)

	if err != nil {
		userCgf.Port = 8089
	}

	jsonErr := json.Unmarshal(configFile, &userCgf)

	if jsonErr != nil {

		userCgf.Port = 8089
	}

	Cfg.userConfig = userCgf

}
