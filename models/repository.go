package models

import (
	"MareWood/helper"
	"MareWood/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
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

var fillWhiteListForRepository = []string{"desc", "name"}

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

type RepositoryPageResult struct {
	List      []*Repository
	Total     int
	PageNum   int
	PageSize  int
	TotalPage int
}

func (r *Repository) UpdateFieldContent(id string, field string, fieldContent string) (err error) {
	field = strings.ToLower(field)

	if !helper.InStrArr(field, fillWhiteListForRepository) {
		return errors.New("不能修改当前字段！")
	}

	return sql.DB.Model(&r).Where("id = ?", id).Update(field, fieldContent).Error
}

func (r *Repository) Create() error {

	if !sql.DB.Where("url = ?", r.Url).First(&Repository{}).RecordNotFound() {
		return errors.New("repository already exists")
	}

	r.Status = RepoStatusProcessing
	r.JobStatus = RepoJobStatusLeisured

	return sql.DB.Create(&r).Error

}

func FindRepository(name string, pageNum int, pageSize int, maps interface{}) (RepositoryPageResult, error) {
	var result RepositoryPageResult
	var err error

	result.List,result.Total, err = GetRepositories(name, pageNum, pageSize, maps)
	if err != nil {
		return result, err
	}

	result.PageSize = pageSize
	result.PageNum = pageNum

	result.TotalPage = helper.ComputeTotalPage(result.Total, pageSize)

	return result, nil
}

func GetRepositories(name string,pageNum int, pageSize int, maps interface{}) ([]*Repository, int, error) {

	var total int
	var repositories []*Repository

	query := sql.DB.Model(&Repository{}).Where(maps).Order("updated_at desc")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+ name +"%")
	}
	query.Count(&total)
	err := query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&repositories).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return repositories, 0, err
	}
	return repositories, total, nil
}
