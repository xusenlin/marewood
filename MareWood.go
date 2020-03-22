package main

import (
	_ "MareWood/config"
	"MareWood/database"
	"MareWood/helper"
	"MareWood/routes"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)


func main() {

	if err := helper.CheckEnvAndInitDir(); err !=nil{
		log.Printf("检查命令工具或者初始化目录失败！")
		log.Printf("请安装git和node以及在.env配置里所有前端依赖安装工具")
		log.Fatal(err)
	}
	database.InitDb()
	defer database.DB.Close()

	r := routes.InitRouter()
	if err := r.Run(":" + os.Getenv("HTTP_PORT")); err != nil {
		panic(err)
	}
}
