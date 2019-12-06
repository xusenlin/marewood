package config

import (
	"FEDeployService/helper"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbDns          string
	Version        float32
	AppName        string
	CurrentDir     string
	AppRepository  string
	RepositoryDir  string
	ResourcesDir   string
	WebRootDir     string
	ClientDir      string
	WebHookUrl     string
	DependTools []string
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
	Cfg.AppName = "FEDeployService"
	Cfg.AppRepository = "https://github.com/xusenlin/FEDeployService"
	Cfg.DbDns = Cfg.CurrentDir + "/database.db" //数据库地址

	Cfg.ResourcesDir = Cfg.CurrentDir + "/resources"       //所有仓库和打包生成的文件
	Cfg.RepositoryDir = Cfg.ResourcesDir + "/repositories" //克隆过来的所有仓库
	Cfg.WebRootDir = Cfg.ResourcesDir + "/webs"            //打包完成后的静态文件存放地址
	Cfg.ClientDir = Cfg.CurrentDir + "/public/build"       //客户端地址
	Cfg.WebHookUrl = "/web_hook"                           //仓库pull事件触发时，会带上仓库id触发此接口做本地仓库更新代码和根据package.json变动跟新依赖，并触发此仓库关联的自动任务
	Cfg.DependTools = helper.GetDependTools()

}
