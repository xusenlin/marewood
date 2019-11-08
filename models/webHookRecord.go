package models

import "github.com/jinzhu/gorm"


const WebHookRecordFail = 2
const WebHookRecordSuccess = 1

type WebHookRecord struct {
	gorm.Model
	Status       int
	TerminalInfo string `gorm:"type:varchar(1000)"`
	RepositoryId uint
}
