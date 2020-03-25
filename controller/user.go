package controller

import (
	"MareWood/database"
	"MareWood/helper"
	"MareWood/models"
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


	if !database.DB.Where("username = ?", newUser.Username).First(&models.User{}).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "用户名已经存在",
		})
		return
	}

	newUser.ID = 0
	newUser.Status = models.UserStatusEnabled
	if newUser.Username == "Admin" {
		newUser.Role =  models.UserRoleSuperAdministrator
	}else {
		newUser.Role =  models.UserRoleReporter
	}
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

	token,err := helper.GenToken(&models.Claims{
		ID:user.ID,
		Username:user.Username,
		Role:user.Role,
		Status:user.Status,
	})

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

func UserFindAll(c *gin.Context)  {

	var users []models.User

	if database.DB.Find(&users).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
		"msg":    "查询成功",
	})
}

func UserDestroy(c *gin.Context) {

	id := c.Query("id")

	if !database.DB.
		Where("id = ? AND role = ?", id ,models.UserRoleSuperAdministrator).
		First(&models.User{}).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "超级管理员不能被删除",
		})
		return
	}

	if database.DB.Where("id = ?", id ).Delete(&models.User{}).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}

func UserRoleEdit(c *gin.Context) {

	id := c.Query("id")
	isUp := c.Query("isUp")
	var user models.User

	if database.DB.First(&user,id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}

	if user.Role == models.UserRoleSuperAdministrator {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "超级管理员无法变更角色",
		})
		return
	}

	if isUp == "1" {
		if user.Role == models.UserRoleAdminister {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "已经是管理员了,不能再升级了",
			})
			return
		}
		user.Role = user.Role - 1
	}else {
		if user.Role == models.UserRoleReporter {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "已经是记者身份了,不能再降级了",
			})
			return
		}
		user.Role = user.Role + 1
	}


	if database.DB.Save(&user).Error != nil {
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
		"msg":    "操作成功",
	})
}