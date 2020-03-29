package sql

import (
	"MareWood/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDb() {
	var err error

	DB, err = gorm.Open("sqlite3", config.Cfg.DbDns)

	if err != nil {
		panic("failed to connect database")
	}

	DB.LogMode(config.Cfg.LogMode)
}
