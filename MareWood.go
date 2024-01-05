package main

import (
	"log"
	"marewood/conf"
	"marewood/initial"
	"marewood/internal/db"
	"marewood/routes"
)

func main() {

	if err := initial.CheckEnvAndInitDir(); err != nil {
		log.Fatal(err)
		return
	}
	if err := db.Connection(conf.DbDns); err != nil {
		log.Fatal(err)
		return
	}
	if err := initial.AutoMigrate(); err != nil {
		log.Fatal(err)
		return
	}

	r := routes.InitRouter()

	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		panic(err)
	}

}
