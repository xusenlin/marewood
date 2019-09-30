package main

import (
	"FrontEndAutomatedDeployment/database"
	"FrontEndAutomatedDeployment/routes"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)


func main() {
	defer database.DB.Close()

	r := routes.InitRouter()

	err := r.Run(":" + os.Getenv("HTTP_PORT")) // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}
}
