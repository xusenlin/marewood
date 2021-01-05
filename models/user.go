package models

import (
	"MareWood/helper"
	"MareWood/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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
	PasswordSalt        = "MareWood"//密码加盐
	TokenExpireDuration = time.Hour * 10 //Token过期时间
	JwtClaimsKey        = "JwtClaims" //Claims 储存在 *gin.Context 里的关键字，通过JWTAuth中间件设置。
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

func (u *User) Register() (err error) {

	if !sql.DB.Where("username = ?", u.Username).First(&User{}).RecordNotFound() {
		return errors.New("username already exists")
	}

	u.ID = 0
	u.Status = UserStatusEnabled
	if u.Username == "Admin" {
		u.Role = UserRoleSuperAdministrator
	} else {
		u.Role = UserRoleReporter
	}
	u.Password = helper.DigestString(PasswordSalt + u.Password)

	return sql.DB.Save(&u).Error
}

func (u *User) Login() (err error) {

	password := helper.DigestString(PasswordSalt + u.Password)

	if sql.DB.
		Where("username = ? AND password = ?", u.Username, password).
		First(&u).RecordNotFound() {
		return errors.New("username or password is wrong")
	}
	if u.Status == UserStatusDisabled {
		return errors.New("user is disabled")
	}

	u.Password = "***"

	return nil
}
