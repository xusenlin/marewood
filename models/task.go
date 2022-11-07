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
	TaskStatusFail       = 2 //打包失败
	TaskStatusSuccess    = 1 //打包成功
	TaskStatusLeisured   = 0 //默认状态，空闲
	TaskStatusProcessing = 3 //正在打包状态
)

var fillWhiteListForTask = []string{"desc", "name"}

type Task struct {
	gorm.Model
	Name          string `binding:"required,min=2,max=20"`
	Desc          string `gorm:"type:varchar(1000)",binding:"required,min=2,max=999"`
	Status        int    `gorm:"default:0"`        //任务状态
	Branch        string `gorm:"default:'main'"` //部署分支默认main，用户在部署之前随时可以修改
	Url           string                           //访问目录
	RunQuantity   int `gorm:"default:0"`
	CategoryId    int `gorm:"index",binding:"required"`
	WebHookUrl    string
	RepositoryId  int    `gorm:"index",binding:"required"`
	BuildDir      string `binding:"required"` //打包的目录,默认是dist
	BuildCommand  string `binding:"required"` //打包命令，npm run build 可以读取package.json供选择
	User          string                      //操作人，目前只有加锁用户
	LockPassword  string                      //任务加锁
	TerminalInfo  string `gorm:"type:varchar(1000)"`
	SuccessScript string `gorm:"type:varchar(1000)"` //打包成功运行的脚本，多个用 ; 隔开
}

type TaskPageResult struct {
	List      []*Task
	Total     int
	PageNum   int
	PageSize  int
	TotalPage int
}

func (t *Task) FindAll() (tasks []Task, err error) {
	err = sql.DB.Order("created_at desc").Select("id,name").Find(&tasks).Error
	return
}

func (t *Task) Create() (err error) {

	t.RunQuantity = 0
	t.Branch = "master"
	t.Status = TaskStatusLeisured
	err = sql.DB.Create(&t).Error
	if err != nil {
		return
	}

	t.WebHookUrl = config.Cfg.WebHookUrl + "?id=" + strconv.Itoa(int(t.ID))
	err = sql.DB.Save(&t).Error

	return
}

func (t *Task) UpdateBranch(branch string) (err error) {
	err =
		sql.DB.Model(&t).Update("branch", branch).Error
	return
}

func (t *Task) UpdateFieldContent(id string, field string, fieldContent string) (err error) {

	field = strings.ToLower(field)

	if !helper.InStrArr(field, fillWhiteListForTask) {
		return errors.New("cannot modify the current field！")
	}

	if err := sql.DB.First(&t, id).Error; err != nil {
		return err
	}
	if t.LockPassword != "" {
		return errors.New("please unlock the task first")
	}

	return sql.DB.Model(&t).Update(field, fieldContent).Error

}

func (t *Task) Destroy(id string) (err error) {
	err = sql.DB.First(&t, id).Error
	if err != nil {
		return
	}
	err = new(Category).CategoryTaskQuantityDecrement(t.CategoryId)
	if err != nil {
		return
	}

	return sql.DB.Delete(&t).Error
}

func FindTask(name string, pageNum int, pageSize int, maps interface{}) (TaskPageResult, error) {
	var result TaskPageResult
	var err error

	result.List, result.Total, err = GetTasks(name,pageNum, pageSize, maps)
	if err != nil {
		return result, err
	}

	result.PageSize = pageSize
	result.PageNum = pageNum

	result.TotalPage = helper.ComputeTotalPage(result.Total, pageSize)

	return result, nil
}

func GetTasks(name string, pageNum int, pageSize int, maps interface{}) ([]*Task, int, error) {
	var total int
	var tasks []*Task

	query := sql.DB.Model(&Task{}).Where(maps).Order("updated_at desc")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	query.Count(&total)
	err := query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tasks).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return tasks, 0, err
	}
	return tasks, total, nil
}
