package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Name         string
	Desc         string `gorm:"type:varchar(1000);not null;default:''"`
	RepositoryId int    `gorm:"index"`
	Branch       string `gorm:"not null;default:'master';comment:'部署分支'"`
	RootDir      string `gorm:"not null;default:'';comment:'此次任务的根目录，默认就是仓库根目录，在一个仓库有多个项目的时候用'"`
	AccessUrl    string `gorm:"not null;default:'';comment:'访问的url,当此次任务对应的目录没有index.html不返回url'"`
	DeployWay    int    `gorm:"not null;default:1;comment:'部署方式，1手动，2自动，3定时'"`
	BuildDir     string `gorm:"not null;default:'dist';comment:'打包目录'"`
	BuildCommand string `gorm:"not null;default:'npm run build';comment:'打包命令，可以读取package.json供选择'"`
	//webhook 签名
}
