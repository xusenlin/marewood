package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string
	Desc string `gorm:"type:varchar(1000);not null;default:''"`
}
