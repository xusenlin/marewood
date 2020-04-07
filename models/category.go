package models

import (
	"MareWood/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Category struct {
	gorm.Model
	Name        string `binding:"required,min=2,max=20"`
	JobQuantity int    `gorm:"default:0"`
	Desc        string `gorm:"type:varchar(1000)"`
}

func (c *Category) CategoryFindAll() (categories []Category, err error) {
	err = sql.DB.Order("created_at asc").Find(&categories).Error
	return
}


func (c *Category) Create() (err error) {
	c.JobQuantity = 0
	err = sql.DB.Create(&c).Error
	return
}
func (c *Category) UpdateDesc(id string, desc string) (err error) {
	err =
		sql.DB.Model(&c).Where("id = ?", id).
			UpdateColumn("desc", desc).Error
	return
}

func (c *Category) CategoryJobQuantityIncrement(id int) (err error) {
	err =
		sql.DB.Model(&c).Where("id = ?", id).
		UpdateColumn("job_quantity", gorm.Expr("job_quantity + ?", 1)).Error
	return
}


func (c *Category) CategoryJobQuantityDecrement(id int) (err error) {
	err =
		sql.DB.Model(&c).Where("id = ?", id).
			UpdateColumn("job_quantity", gorm.Expr("job_quantity - ?", 1)).Error
	return
}


func (c *Category) Destroy(id string) (err error) {

	var jobCount int

	err = sql.DB.Model(&Job{}).
		Where("category_id = ?", id).
		Count(&jobCount).Error

	if err != nil {
		return err
	}
	if jobCount > 0 {
		return errors.New("无法删除，还有" + strconv.Itoa(jobCount) + "个任务在使用此分类")
	}
	return sql.DB.Where("id = ?", id ).Delete(&c).Error
}
