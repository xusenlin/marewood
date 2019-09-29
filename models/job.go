package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Name          string
	Desc          string `gorm:"type:varchar(1000)"`
	RepositoryId  int    `gorm:"index"`
	Branch        string                          //部署分支master
	RootDir       string                          //此次任务的根目录，默认就是仓库根目录，在一个仓库有多个项目的时候用
	AccessUrl     string                          //访问的url,当此次任务对应的目录没有index.html不返回url
	DeployWay     int `gorm:"not null;default:0"` //部署方式，1手动，2自动，3定时
	BuildDir      string                          //打包目录dist
	BuildCommand  string                          //打包命令，npm run build 可以读取package.json供选择
	WebHookSecret string                          //webHook密钥
	WebHookUrl    string                          //webHook 地址
}
