package main

import (
	"FEDeployService/database"
	"FEDeployService/routes"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)


func main() {
	defer database.DB.Close()

	r := routes.InitRouter()

	err := r.Run(":" + os.Getenv("HTTP_PORT"))

	if err != nil {
		panic(err)
	}
}
