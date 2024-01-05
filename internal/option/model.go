package option

import (
	"marewood/internal/common"
	"marewood/internal/db"
)

const KeyWebUrl = "StaticWebUrl"

type Option struct {
	common.Model
	Key   string `gorm:"type:varchar(45);unique;" binding:"required,min=2,max=45" json:"key"`
	Value string `gorm:"type:varchar(1000);" json:"value"`
}

func (o *Option) Update() (err error) {
	return db.Conn.Save(&o).Error
}

func (o *Option) Create() (err error) {
	return db.Conn.Create(&o).Error
}

func (o *Option) Get() (err error) {
	return db.Conn.Where("key = ?", o.Key).First(&o).Error
}
