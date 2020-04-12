package models

import (
	"MareWood/sql"
	"errors"
	"github.com/jinzhu/gorm"
)

const ( //仓库拉取状态
	RepoStatusFail       = 2
	RepoStatusSuccess    = 1
	RepoStatusProcessing = 0
)
const ( //仓库工作状态
	RepoJobStatusBusy     = 1
	RepoJobStatusLeisured = 0
)

type Repository struct {
	gorm.Model
	Name         string `binding:"required,min=2,max=20"`
	Desc         string `gorm:"type:varchar(1000)"`
	Url          string `binding:"url"`      //仓库的地址
	UserName     string                      //仓库是私有的话需要填写
	Password     string                      //仓库是私有的话需要填写
	Status       int                         //此仓库是否已经将代码克隆过来
	JobStatus    int                         //有任务正在打包，其他任务无法执行
	DependTools  string `binding:"required"` //依赖工具选择
	TerminalInfo string `gorm:"type:varchar(1000)"`
}

func (r *Repository) UpdateDesc(id string, desc string) (err error) {
	err =
		sql.DB.Model(&r).Where("id = ?", id).
			UpdateColumn("desc", desc).Error
	return
}

func (r *Repository) Create() error {

	if !sql.DB.Where("url = ?", r.Url).First(&Repository{}).RecordNotFound() {
		return errors.New("仓库已经存在")
	}

	r.Status = RepoStatusProcessing
	r.JobStatus = RepoJobStatusLeisured

	return sql.DB.Create(&r).Error

}
