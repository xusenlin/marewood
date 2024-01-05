package db

import (
	ginPagination "github.com/xusenlin/gin-pagination"
	//"gorm.io/driver/sqlite"  Remove cgo to use a slim Alpine Docker container
	//Because of C-compiler requirement, you can't build your Go code inside tiny stripped containers like (golang-alpine)
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connection(dns string) error {

	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}
	//Conn = db.Debug()
	Conn = db
	ginPagination.Init(&ginPagination.Config{
		PageSizeMaxVal:     100,
		PageSizeDefaultVal: 20,
		DB:                 Conn,
	})
	return nil
}
