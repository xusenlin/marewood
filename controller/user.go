package controller

import (
	"FEDeployService/database"
	"FEDeployService/helper"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	newUser.ID = 0
	newUser.Role = models.UserRoleReporter
	newUser.Status = models.UserStatusEnabled
	newUser.Password = helper.DigestString(models.PasswordSalt + newUser.Password)

	if database.DB.Save(&newUser).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "注册成功",
	})

}

func UserLogin(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

 	password := helper.DigestString(models.PasswordSalt + user.Password)


	if database.DB.
		Where("username = ? AND password = ?", user.Username, password).
 		First(&user).RecordNotFound() {

		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "用户名或者密码出错",
		})
		return
	}
 	user.Password = "***"

	token,err := helper.GenToken(user.ID,user.Username,user.Role,user.Status)

	if  err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	type userAllInfo struct {
		models.User
		Token string
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   userAllInfo{user, token,},
		"msg":    "登录成功",
	})

}