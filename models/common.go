package models

import "MareWood/sql"

//type SysSetting struct {
//	gorm.Model
//	OptionKey   string `binding:"required"`
//	OptionValue string `gorm:"type:varchar(3000)",binding:"required"`
//}

type PackageJson struct {
	Name  string
	Version string
	Private bool
	Dependencies interface{}
	DevDependencies interface{}
	Scripts interface{}
}

func AutoMigrate()  {
	sql.DB.AutoMigrate(
		&Job{},
		&User{},
		&Category{},
		&Repository{},
		//&models.SysSetting{},
	)
}