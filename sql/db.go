package sql

import (
	"MareWood/config"
	"MareWood/models"
	"fmt"
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

	DB.AutoMigrate(
		&models.Job{},
		&models.User{},
		&models.Category{},
		&models.Repository{},
		//&models.SysSetting{},
	)

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
