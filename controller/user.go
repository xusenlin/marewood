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
		"msg":    "registration success",
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
		"msg":    "login successful",
	})

}

func UserFindAll(c *gin.Context) {

	var users []models.User
	keyword := c.Query("name")

	query := sql.DB.Order("created_at asc")

	if keyword != "" {
		query = query.Where("username LIKE ?", "%"+ keyword +"%")
	}

	if query.Find(&users).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
			"msg":    "Database query error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
		"msg":    "search successful",
	})
}

func UserDestroy(c *gin.Context) {

	id := c.Query("id")
	var user  models.User

	if !sql.DB.
		Where("id = ? AND role = ?", id, models.UserRoleSuperAdministrator).
		First(&models.User{}).RecordNotFound() {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "Super administrator cannot be deleted",
		})
		return
	}

	if sql.DB.Where("id = ?", id).Delete(&user).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	claims, _ := serviceUser.GetJwtClaimsByContext(c)
	msg := models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” delete the user“" + user.Username + "”",
	}
	models.Broadcast <- msg

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
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
			"msg":    "Super administrator cannot change role",
		})
		return
	}

	if isUp == "1" {
		if user.Role == models.UserRoleAdminister {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "Already an administrator, can no longer upgrade",
			})
			return
		}
		user.Role = user.Role - 1
	} else {
		if user.Role == models.UserRoleReporter {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"data":   "",
				"msg":    "Already a reporter, cannot be demoted",
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
		"msg":    "Successful operation",
	})
}
