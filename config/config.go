package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbDns         string
	Version       float32
	AppName       string
	CurrentDir    string
	AppRepository string
	RepositoryDir string
	WebRootDir    string
	ClientDir     string
}

var Cfg Config

func init() {
	var err error
	if Cfg.CurrentDir, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg.Version = 0.1
	Cfg.AppName = "FEAutoDeploy"
	Cfg.AppRepository = "https://github.com/xusenlin/FrontEndAutomatedDeployment"
	Cfg.DbDns = Cfg.CurrentDir + "/database.db"
	Cfg.RepositoryDir = Cfg.CurrentDir + "/resources/repository" //克隆过来的所有仓库
	Cfg.WebRootDir = Cfg.CurrentDir + "/resources/webs"          //打包完成后的静态文件存放地址
	Cfg.ClientDir = Cfg.CurrentDir + "/public/build"             //客户端地址

}
