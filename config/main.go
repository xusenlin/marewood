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

	systemCfg.Version = 0.2
	systemCfg.AppName = "MareWood"
	systemCfg.AppRepository = "https://github.com/xusenlin/MareWood"
	systemCfg.DbDns = systemCfg.CurrentDir + "/database.db"
	systemCfg.ResourcesDir = systemCfg.CurrentDir + "/resources"
	systemCfg.RepositoryDir = systemCfg.ResourcesDir + "/repositories"
	systemCfg.WebRootDir = systemCfg.ResourcesDir + "/webs"
	systemCfg.ClientDir = ClientDir
	systemCfg.WebHookUrl = "/web_hook"
	systemCfg.WebsUrl = "/webs"
	systemCfg.Introduction = systemCfg.AppName + ` 是一个轻量级的前端部署工具，使用了 GOLANG、GIN、GORM、JWT、SQLITE、REACE、MATERIAL-UI 开发，
不同于 Jenkins 的大而全，它很简单且只针对前端，你可以很灵活的配置各种部署环境。
如果你愿意，线上发布也可以是点击一下按钮这么简单的事情,当然也可以配置 WEBHOOK，提交 GIT 代码既自动发布。
`
	Cfg.systemConfig = systemCfg
	Cfg.userConfig = userCfg
}

func GetSysInfo() []interface{} {
	type rowInfo struct {
		Title string
		Val   interface{}
		Desc  string
	}
	return []interface{}{
		rowInfo{"名字", Cfg.AppName, "系统名字"},
		rowInfo{"版本号", Cfg.Version, "系统版本"},
		rowInfo{"软件目录", Cfg.CurrentDir, "软件运行的目录"},
		rowInfo{"数据库", Cfg.DbDns, "数据库地址（DbDns）"},
		rowInfo{"资源目录", Cfg.ResourcesDir, "包含仓库和打包完成的所有资源"},
		rowInfo{"仓库目录", Cfg.RepositoryDir, "克隆过来的所有仓库目录"},
		rowInfo{"前端打包目录", Cfg.WebRootDir, "所有任务执行成功都会将打包资源存放在这里"},
		rowInfo{"客户端地址", Cfg.ClientDir, "当前系统的前端页面存放地址"},
		rowInfo{"依赖工具", Cfg.DependTools, "当前服务器支持的前端依赖安装工具"},
		rowInfo{"WebHookUrl", Cfg.WebHookUrl, "用于触发任务的WebHook地址，值得注意的是，正在繁忙的任务和加密任务是不会执行的"},
		rowInfo{"WebsUrl", Cfg.WebsUrl, "打包成功访问页面的地址"},
		rowInfo{"开源地址", Cfg.AppRepository, "开源的GITHUB地址"},
		rowInfo{"简介", Cfg.Introduction, "应用简介"},
	}
}
