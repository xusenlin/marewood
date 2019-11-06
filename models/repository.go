package models

import "github.com/jinzhu/gorm"

type Repository struct {
	gorm.Model
	Name          string
	Desc          string `gorm:"type:varchar(1000)"`
	Url           string //仓库的地址
	WebHookSecret string //webHook密钥
	UserName      string //仓库是私有的话需要填写
	Password      string //仓库是私有的话需要填写
	//WebHookUrl    string //webHook 地址 这里不再需要通过公共的接口根据仓库id来
	//仓库存放在
}
