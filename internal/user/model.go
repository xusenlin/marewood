package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"marewood/internal/common"
	"marewood/internal/db"
)

func Md5(p string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p)))
}

type User struct {
	common.Model
	Username    string `gorm:"unique" binding:"required,min=2,max=8" json:"username"`
	Password    string `json:"password"`
	Role        int    `json:"role"`
	Status      int    `gorm:"default:1" json:"status"`
	Avatar      string `json:"avatar"`
	Description string `gorm:"type:varchar(200)" json:"description"`
}

func (u *User) Register() (err error) {

	if !errors.Is(db.Conn.Where("username = ?", u.Username).First(&User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("username already exists")
	}
	u.Status = StatusEnabled
	if u.Username == "Admin" {
		u.Role = RoleSuperAdministrator
	} else {
		u.Role = RoleReporter
	}
	u.Password = Md5(PasswordSalt + u.Password)

	return db.Conn.Save(&u).Error
}

func (u *User) Login() (err error) {

	password := Md5(PasswordSalt + u.Password)

	if errors.Is(db.Conn.
		Where("username = ? AND password = ?", u.Username, password).
		First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("username or password is wrong")
	}

	if u.Status == StatusDisabled {
		return errors.New("user is disabled")
	}

	u.Password = "***"

	return nil
}

func (u *User) Destroy() (err error) {
	if errors.Is(db.Conn.First(&u).Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	if u.Role == RoleSuperAdministrator {
		return errors.New("can not delete super administrator")
	}
	return db.Conn.Delete(&u).Error
}

func (u *User) Edit() (err error) {

	password := u.Password

	if u.Role == RoleSuperAdministrator {
		return errors.New("can not edit super administrator")
	}
	if password != "" {
		u.Password = Md5(PasswordSalt + password)
	}
	return db.Conn.Save(&u).Error
}
