package dao

import (
	"gorm.io/gorm"
	"marewood/internal/models"
	"marewood/internal/pkg/errors"
)

// RepositoryDAO 仓库数据访问接口
type RepositoryDAO interface {
	Create(repo *models.Repository) error
	FindByID(id uint) (*models.Repository, error)
	FindByURL(url string, excludeID uint) (*models.Repository, error)
	Update(repo *models.Repository) error
	Delete(id uint) error
	List(offset, limit int, filters map[string]interface{}) ([]*models.Repository, int64, error)
	ListAll() ([]*models.Repository, error)
	CountTasksByRepositoryID(repoID uint) (int64, error)
}

type repositoryDAO struct {
	db *gorm.DB
}

// NewRepositoryDAO 创建仓库 DAO
func NewRepositoryDAO(db *gorm.DB) RepositoryDAO {
	return &repositoryDAO{db: db}
}

func (d *repositoryDAO) Create(repo *models.Repository) error {
	return d.db.Create(repo).Error
}

func (d *repositoryDAO) FindByID(id uint) (*models.Repository, error) {
	var repo models.Repository
	err := d.db.First(&repo, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("repository")
	}
	return &repo, err
}

func (d *repositoryDAO) FindByURL(url string, excludeID uint) (*models.Repository, error) {
	var repo models.Repository
	query := d.db.Where("url = ?", url)
	if excludeID != 0 {
		query = query.Where("id != ?", excludeID)
	}
	err := query.First(&repo).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &repo, err
}

func (d *repositoryDAO) Update(repo *models.Repository) error {
	return d.db.Save(repo).Error
}

func (d *repositoryDAO) Delete(id uint) error {
	return d.db.Delete(&models.Repository{}, id).Error
}

func (d *repositoryDAO) List(offset, limit int, filters map[string]interface{}) ([]*models.Repository, int64, error) {
	var repos []*models.Repository
	var total int64

	query := d.db.Model(&models.Repository{})

	// 应用过滤条件
	for key, value := range filters {
		if key == "name" {
			query = query.Where("name LIKE ?", "%"+value.(string)+"%")
		} else {
			query = query.Where(key+" = ?", value)
		}
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	err := query.Offset(offset).Limit(limit).Find(&repos).Error
	return repos, total, err
}

func (d *repositoryDAO) ListAll() ([]*models.Repository, error) {
	var repos []*models.Repository
	err := d.db.Select("id", "name", "url").Find(&repos).Error
	return repos, err
}

func (d *repositoryDAO) CountTasksByRepositoryID(repoID uint) (int64, error) {
	var count int64
	err := d.db.Table("tasks").Where("repository_id = ?", repoID).Count(&count).Error
	return count, err
}
