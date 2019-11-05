package database

import (
	"FEDeployService/config"
	"FEDeployService/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init()  {
	var err error

	DB, err = gorm.Open("sqlite3", config.Cfg.DbDns)

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.SysSetting{},&models.Category{},&models.Repository{},&models.Job{})

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
