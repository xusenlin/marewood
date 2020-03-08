package models

import "github.com/jinzhu/gorm"

const ( //仓库拉取状态
	JobStatusFail       = 2
	JobStatusSuccess    = 1
	JobStatusProcessing = 0
)


type Job struct {
	gorm.Model
	Name         string `binding:"required,min=2,max=16"`
	Desc         string `gorm:"type:varchar(1000)"`
	Status       int                         //任务状态
	Branch       string `binding:"required"` //部署分支master，用户在部署之前随时可以修改
	AccessDir    string                      //访问目录，只有状态成功才返回
	CategoryId   int `gorm:"index"`          // 外键 (属于), tag `index`是为该列创建索引
	RepositoryId int `gorm:"index",binding:"required"`
	BuildDir     string //打包的目录,默认是dist
	BuildCommand string //打包命令，npm run build 可以读取package.json供选择
}
