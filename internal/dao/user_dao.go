package dao

import (
	"gorm.io/gorm"
	"marewood/internal/models"
	"marewood/internal/pkg/errors"
)

// UserDAO 用户数据访问接口
type UserDAO interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	List(offset, limit int, filters map[string]interface{}) ([]*models.User, int64, error)
}

type userDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建用户 DAO
func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{db: db}
}

func (d *userDAO) Create(user *models.User) error {
	return d.db.Create(user).Error
}

func (d *userDAO) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := d.db.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("user")
	}
	return &user, err
}

func (d *userDAO) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("user")
	}
	return &user, err
}

func (d *userDAO) Update(user *models.User) error {
	return d.db.Save(user).Error
}

func (d *userDAO) Delete(id uint) error {
	return d.db.Delete(&models.User{}, id).Error
}

func (d *userDAO) List(offset, limit int, filters map[string]interface{}) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	query := d.db.Model(&models.User{})

	// 应用过滤条件
	for key, value := range filters {
		if key == "username" {
			query = query.Where("username LIKE ?", "%"+value.(string)+"%")
		} else {
			query = query.Where(key+" = ?", value)
		}
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	err := query.Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}
