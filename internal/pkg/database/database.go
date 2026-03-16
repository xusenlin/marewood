package database

import (
	"log/slog"
	"marewood/internal/pkg/logger"

	ginPagination "github.com/xusenlin/gin-pagination"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

// Connection 连接数据库
func Connection(dns string) error {
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}
	//Conn = db.Debug()
	Conn = db
	logger.Slog.Info("Connect to database successfully", slog.String("dns", dns))
	ginPagination.Init(&ginPagination.Config{
		PageSizeMaxVal:     100,
		PageSizeDefaultVal: 20,
		DB:                 Conn,
	})
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return Conn
}
