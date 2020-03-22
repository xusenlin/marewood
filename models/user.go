package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

const ( //用户状态
	UserStatusEnabled  = 1
	UserStatusDisabled = 2
)

const ( //用户角色
	UserRoleReporter           = 4
	UserRoleDeveloper          = 3
	UserRoleAdminister         = 2
	UserRoleSuperAdministrator = 1
)

const (
	PasswordSalt        = "MareWood"
	TokenExpireDuration = time.Hour * 3
)

var TokenSecret = []byte("MareWood")

type Claims struct {
	ID       uint
	Username string
	Role     int
	Status   int
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Username string `gorm:"unique",binding:"required,min=2,max=8"`
	Password string `binding:"required,min=6,max=16"`
	Role     int
	Status   int `gorm:"default:1"`
	Avatar   string
	Desc     string `gorm:"type:varchar(200)"`
}
