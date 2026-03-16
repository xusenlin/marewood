package main

import (
	l "log"
	"marewood/conf"
	"marewood/initial"
	"marewood/internal/pkg/database"
	"marewood/internal/pkg/logger"
	"marewood/routes"
	"path/filepath"
)

func main() {

	if err := initial.CheckEnvAndInitDir(); err != nil {
		l.Fatal(err)
		return
	}
	if err := logger.InitLogger(filepath.Join(conf.ResourcesDir, "log")); err != nil {
		l.Fatal(err)
		return
	}
	if err := database.Connection(conf.DbDns); err != nil {
		logger.Slog.Error(err.Error())
		return
	}
	if err := initial.AutoMigrate(); err != nil {
		logger.Slog.Error(err.Error())
		return
	}

	r := routes.InitRouter()

	logger.Slog.Info("server start at port: " + conf.HttpPort)
	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		logger.Slog.Error(err.Error())
		panic(err)
	}

}
