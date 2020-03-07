package models

import "github.com/jinzhu/gorm"

//随机生成的密钥长度
const RepoWebHookSecretRandSeqLen = 8

type Job struct {
	gorm.Model
	Name         string `binding:"required,min=2,max=16"`
	Desc         string `gorm:"type:varchar(1000)"`
	RepositoryId int    `gorm:"index",binding:"required"`
	Branch       string `binding:"required"`     //部署分支master
	RootDir      string                          //此次任务的根目录，默认就是仓库根目录，在一个仓库有多个项目的时候用
	AccessDir    string                          //复制dist到那个目录下提供访问,当此次任务对应的目录没有index.html不返回url
	CategoryId   int `gorm:"index"`              // 外键 (属于), tag `index`是为该列创建索引
	DeployWay    int `gorm:"not null;default:0"` //部署方式，1手动，2自动，3定时
	BuildDir     string                          //打包目录,默认是dist
	BuildCommand string                          //打包命令，npm run build 可以读取package.json供选择
}
