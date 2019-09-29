package models

import "github.com/jinzhu/gorm"

type Repository struct {
	gorm.Model
	Name       string
	CategoryId int    `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
	Desc       string `gorm:"type:varchar(1000)"`
	Url        string
}
