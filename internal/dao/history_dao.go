package dao

import (
	"marewood/internal/models"

	"gorm.io/gorm"
)

type HistoryDao interface {
	FindByTaskID(taskId uint) ([]models.History, error)
	FindByID(id uint) (*models.History, error)
	Create(history *models.History) error
	Update(history *models.History) error
	FindByTaskIDAndGitHash(taskId uint, gitHash string) (*models.History, error)
}

type historyDao struct {
	db *gorm.DB
}

func NewHistoryDao(db *gorm.DB) HistoryDao {
	return &historyDao{db: db}
}

func (d *historyDao) FindByTaskID(taskId uint) ([]models.History, error) {
	var histories []models.History
	err := d.db.Where("task_id = ?", taskId).Order("created_at desc").Limit(10).Find(&histories).Error
	return histories, err
}

func (d *historyDao) FindByID(id uint) (*models.History, error) {
	var history models.History
	err := d.db.First(&history, id).Error
	return &history, err
}

func (d *historyDao) Create(history *models.History) error {
	return d.db.Create(history).Error
}

func (d *historyDao) Update(history *models.History) error {
	return d.db.Save(history).Error
}

func (d *historyDao) FindByTaskIDAndGitHash(taskId uint, gitHash string) (*models.History, error) {
	var history models.History
	err := d.db.Where("task_id = ? AND git_hash = ?", taskId, gitHash).First(&history).Error
	return &history, err
}
