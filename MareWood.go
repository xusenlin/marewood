package main

import (
	"MareWood/config"
	"MareWood/models"
	"MareWood/routes"
	"MareWood/service/serviceConfig"
	"MareWood/sql"
	"log"
)

func main() {

	err := serviceConfig.CheckEnvAndInitDir()
	if err != nil {
		log.Printf("检查命令工具或者初始化目录失败！")
		log.Printf("请安装git和node以及在.env配置里所有前端依赖安装工具")
		log.Fatal(err)
	}

	sql.InitDb()
	defer sql.DB.Close()

	models.AutoMigrate()
	go models.BroadcastMessages()

	r := routes.InitRouter()

	err = r.Run(":" + config.Cfg.HttpPort)
	if err != nil {
		panic(err)
	}
}
