package main

import (
	"AutomatedDeployment/config"
	"AutomatedDeployment/database"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func main() {
	defer database.DB.Close()
	fmt.Println(config.Cfg.Port)
}
