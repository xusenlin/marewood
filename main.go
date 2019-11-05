package main

import (
	_ "FEDeployService/config"
	"FEDeployService/database"
	"FEDeployService/routes"
	"FEDeployService/service"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)


func main() {

	if err := service.CheckEnvAndInitDir(); err !=nil{
		panic(err)
	}
	database.InitDb()
	defer database.DB.Close()

	r := routes.InitRouter()


	if err := r.Run(":" + os.Getenv("HTTP_PORT")); err != nil {
		panic(err)
	}
}
