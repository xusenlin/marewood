package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
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
	Cfg.DbDns = Cfg.CurrentDir + "/database.db"

	Cfg.ResourcesDir = Cfg.CurrentDir + "/resources"
	Cfg.RepositoryDir = Cfg.ResourcesDir + "/repositories"
	Cfg.WebRootDir = Cfg.ResourcesDir + "/webs"
	Cfg.ClientDir = Cfg.CurrentDir + "/public/build"
	Cfg.WebHookUrl = "/web_hook"
	Cfg.DependTools = getDependTools()

}


func getDependTools() []string {
	toolsStr := os.Getenv("DEPENDENT_TOOLS")

	if !strings.Contains(toolsStr, "|") {
		return []string{"npm"}
	}
	tools := strings.Split(toolsStr, "|")

	return tools
}


func GetSysInfo() (Config , []interface{}) {
	type rowInfo struct {
		Title string
		Val interface{}
		Desc string
	}
	return Cfg, []interface{}{
		rowInfo{"名字",Cfg.AppName,"系统名字"},
		rowInfo{"版本号",Cfg.Version,"系统版本"},
		rowInfo{"软件目录",Cfg.CurrentDir,"软件运行的目录"},
		rowInfo{"数据库",Cfg.DbDns,"数据库地址（DbDns）"},
		rowInfo{"资源目录",Cfg.ResourcesDir,"包含仓库和打包完成的所有资源"},
		rowInfo{"仓库目录",Cfg.RepositoryDir,"克隆过来的所有仓库目录"},
		rowInfo{"前端打包目录",Cfg.WebRootDir,"执行任务打包成功的目录"},
		rowInfo{"客户端地址",Cfg.ClientDir,"当前系统的前端页面存放地址"},
		rowInfo{"依赖工具",Cfg.DependTools,"当前服务器支持的前端依赖安装工具"},
		rowInfo{"WebHookUrl",Cfg.WebHookUrl,"用于触发任务的WebHook地址"},
		rowInfo{"GITHUB",Cfg.AppRepository,"开源地址"},
	}
}