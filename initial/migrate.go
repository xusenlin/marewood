package initial

import (
	"marewood/internal/models"
	"marewood/internal/pkg/database"
)

func AutoMigrate() error {
	return database.Conn.AutoMigrate(
		&models.User{},
		&models.Repository{},
		&models.Task{},
	)
}
