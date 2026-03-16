package dao

import (
	"marewood/internal/models"
	"marewood/internal/pkg/errors"
	"strings"

	"gorm.io/gorm"
)

// TaskDAO 任务数据访问接口
type TaskDAO interface {
	Create(task *models.Task) error
	FindByID(id uint) (*models.Task, error)
	FindByAlias(alias string, excludeID uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
	List(offset, limit int, filters map[string]interface{}, userID uint, isAdmin bool) ([]*models.Task, int64, error)
	GetAllTags() ([]string, error)
}

type taskDAO struct {
	db *gorm.DB
}

// NewTaskDAO 创建任务 DAO
func NewTaskDAO(db *gorm.DB) TaskDAO {
	return &taskDAO{db: db}
}

func (d *taskDAO) Create(task *models.Task) error {
	return d.db.Create(task).Error
}

func (d *taskDAO) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := d.db.First(&task, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("task")
	}
	return &task, err
}

func (d *taskDAO) FindByAlias(alias string, excludeID uint) (*models.Task, error) {
	var task models.Task
	query := d.db.Where("alias = ?", alias)
	if excludeID != 0 {
		query = query.Where("id != ?", excludeID)
	}
	err := query.First(&task).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &task, err
}

func (d *taskDAO) Update(task *models.Task) error {
	return d.db.Save(task).Error
}

func (d *taskDAO) Delete(id uint) error {
	return d.db.Delete(&models.Task{}, id).Error
}

func (d *taskDAO) List(offset, limit int, filters map[string]interface{}, userID uint, isAdmin bool) ([]*models.Task, int64, error) {
	var tasks []*models.Task
	var total int64

	query := d.db.Model(&models.Task{})

	// 应用过滤条件
	for key, value := range filters {
		if key == "name" {
			query = query.Where("name LIKE ?", "%"+value.(string)+"%")
		} else if key == "tags" {
			tags := strings.Split(value.(string), ",")
			query = query.Where("tag IN ?", tags)
		} else {
			query = query.Where(key+" = ?", value)
		}
	}

	// 权限过滤：非管理员只能看到公开任务或自己的私有任务
	if !isAdmin {
		query = query.Where("(private = ? AND created_id = ?) OR private = ?", true, userID, false)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	err := query.Offset(offset).Limit(limit).Find(&tasks).Error
	return tasks, total, err
}

func (d *taskDAO) GetAllTags() ([]string, error) {
	var tags []string
	err := d.db.Model(&models.Task{}).Distinct().Pluck("tag", &tags).Error
	return tags, err
}
