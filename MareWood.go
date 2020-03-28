package main

import (
	"MareWood/config"
	"MareWood/helper"
	"MareWood/routes"
	"MareWood/sql"
	"log"
)


func main() {

	if err := helper.CheckEnvAndInitDir(); err !=nil{
		log.Printf("检查命令工具或者初始化目录失败！")
		log.Printf("请安装git和node以及在.env配置里所有前端依赖安装工具")
		log.Fatal(err)
	}


	sql.InitDb()
	defer sql.DB.Close()

	r := routes.InitRouter()
	if err := r.Run(":" + config.Cfg.HttpPort ); err != nil {
		panic(err)
	}
}
