package models

import "github.com/jinzhu/gorm"

type SysSetting struct {
	gorm.Model
	OptionKey   string
	OptionValue string `gorm:"type:varchar(3000)"`
}
