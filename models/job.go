package models

import (
	"MareWood/config"
	"MareWood/helper"
	"MareWood/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

const ( //任务执行状态
	JobStatusFail       = 2 //打包失败
	JobStatusSuccess    = 1 //打包成功
	JobStatusLeisured   = 0 //默认状态，空闲
	JobStatusProcessing = 3 //正在打包状态
)

var fillWhiteListForJob = []string{"desc", "name"}

type Job struct {
	gorm.Model
	Name          string `binding:"required,min=2,max=20"`
	Desc          string `gorm:"type:varchar(1000)",binding:"required,min=2,max=999"`
	Status        int    `gorm:"default:0"`        //任务状态
	Branch        string `gorm:"default:'master'"` //部署分支默认master，用户在部署之前随时可以修改
	Url           string //访问目录
	RunQuantity   int    `gorm:"default:0"`
	CategoryId    int    `gorm:"index",binding:"required"`
	WebHookUrl    string
	RepositoryId  int    `gorm:"index",binding:"required"`
	BuildDir      string `binding:"required"` //打包的目录,默认是dist
	BuildCommand  string `binding:"required"` //打包命令，npm run build 可以读取package.json供选择
	User          string //操作人，目前只有加锁用户
	LockPassword  string //任务加锁
	TerminalInfo  string `gorm:"type:varchar(1000)"`
	SuccessScript string `gorm:"type:varchar(1000)"` //打包成功运行的脚本，多个用 ; 隔开
}

type JobPageResult struct {
	List      []*Job
	Total     int
	PageNum   int
	PageSize  int
	TotalPage int
}


func (j *Job) FindAll() (jobs []Job, err error) {
	err = sql.DB.Order("created_at desc").Select("id,name").Find(&jobs).Error
	return
}

func (j *Job) Create() (err error) {

	j.RunQuantity = 0
	j.Branch = "master"
	j.Status = JobStatusLeisured
	err = sql.DB.Create(&j).Error
	if err != nil {
		return
	}

	j.WebHookUrl = config.Cfg.WebHookUrl + "?id=" + strconv.Itoa(int(j.ID))
	err = sql.DB.Save(&j).Error

	return
}

func (j *Job) UpdateBranch(branch string) (err error) {
	err =
		sql.DB.Model(&j).Update("branch", branch).Error
	return
}

func (j *Job) UpdateFieldContent(id string, field string, fieldContent string) (err error) {

	field = strings.ToLower(field)

	if !helper.InStrArr(field, fillWhiteListForJob) {
		return errors.New("不能修改当前字段！")
	}

	if err := sql.DB.First(&j, id).Error; err != nil {
		return err
	}
	if j.LockPassword != "" {
		return errors.New("请先解锁任务")
	}

	return sql.DB.Model(&j).Update(field, fieldContent).Error

}

func (j *Job) Destroy(id string) (err error) {
	err = sql.DB.First(&j, id).Error
	if err != nil {
		return
	}
	err = new(Category).CategoryJobQuantityDecrement(j.CategoryId)
	if err != nil {
		return
	}

	return sql.DB.Delete(&j).Error
}

func FindJob(pageNum int, pageSize int, maps interface{}) (JobPageResult, error) {
	var result JobPageResult
	var err error

	result.List, result.Total, err = GetJobs(pageNum, pageSize, maps)
	if err != nil {
		return result, err
	}

	result.PageSize = pageSize
	result.PageNum = pageNum

	result.TotalPage = helper.ComputeTotalPage(result.Total, pageSize)

	return result, nil
}

func GetJobs(pageNum int, pageSize int, maps interface{}) ([]*Job, int, error) {
	var total int
	var jobs []*Job

	query := sql.DB.Model(&Job{}).Where(maps).Order("updated_at desc")
	query.Count(&total)
	err := query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&jobs).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return jobs, 0, err
	}
	return jobs, total, nil
}
