package main

import (
	l "log"
	"marewood/conf"
	"marewood/initial"
	"marewood/internal/db"
	"marewood/internal/log"
	"marewood/routes"
	"path/filepath"
)

func main() {

	if err := initial.CheckEnvAndInitDir(); err != nil {
		l.Fatal(err)
		return
	}
	if err := log.InitLogger(filepath.Join(conf.ResourcesDir, "log")); err != nil {
		l.Fatal(err)
		return
	}
	if err := db.Connection(conf.DbDns); err != nil {
		log.Slog.Error(err.Error())
		return
	}
	if err := initial.AutoMigrate(); err != nil {
		log.Slog.Error(err.Error())
		return
	}

	r := routes.InitRouter()

	log.Slog.Info("server start at port: " + conf.HttpPort)
	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		log.Slog.Error(err.Error())
		panic(err)
	}

}
