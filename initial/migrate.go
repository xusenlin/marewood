package initial

import (
	"marewood/internal/db"
	"marewood/internal/repository"
	"marewood/internal/task"
	"marewood/internal/user"
)

func AutoMigrate() error {
	return db.Conn.AutoMigrate(
		&task.Task{},
		&user.User{},
		&repository.Repository{},
	)
}
