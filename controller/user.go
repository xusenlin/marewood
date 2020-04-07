package controller

import (
	"MareWood/models"
	"MareWood/service/serviceUser"
	"MareWood/sql"
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

	if err := newUser.Register(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
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

	if err := user.Login(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	token, err := serviceUser.GenToken(&models.Claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	})

	if err != nil {
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

func UserFindAll(c *gin.Context) {

	var users []models.User

	if sql.DB.Find(&users).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
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

	if !sql.DB.
		Where("id = ? AND role = ?", id, models.UserRoleSuperAdministrator).
		First(&models.User{}).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "超级管理员不能被删除",
		})
		return
	}

	if sql.DB.Where("id = ?", id).Delete(&models.User{}).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
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

	if sql.DB.First(&user, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
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
	} else {
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

	if sql.DB.Save(&user).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "操作成功",
	})
}
