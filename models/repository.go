package models

import "github.com/jinzhu/gorm"

const ( //仓库拉取状态
	RepoStatusFail       = 2
	RepoStatusSuccess    = 1
	RepoStatusProcessing = 0
)

const ( //仓库依赖更新情况
	RepoDependStatusFail       = 2
	RepoDependStatusSuccess    = 1
	RepoDependStatusProcessing = 0
)

//随机生成的密钥长度
const RepoWebHookSecretRandSeqLen = 8

type Repository struct {
	gorm.Model
	Name               string
	Desc               string `gorm:"type:varchar(1000)"`
	Url                string //仓库的地址
	WebHookSecret      string //webHook密钥
	UserName           string //仓库是私有的话需要填写
	Password           string //仓库是私有的话需要填写
	Status             int    //此仓库是否已经将代码克隆过来
	DependentTools     string  //依赖工具选择
	TerminalInfo       string `gorm:"type:varchar(1000)"`
	DependStatus       int //此仓库依赖的状态
	DependTerminalInfo string `gorm:"type:varchar(1000)"`
}
