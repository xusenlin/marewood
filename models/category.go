package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string `binding:"required,min=2,max=16"`
	JobQuantity int    `gorm:"default:0"`
	Desc        string `gorm:"type:varchar(1000)"`
}
