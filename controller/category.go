package controller

import (
	"MareWood/models"
	"MareWood/service/serviceUser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindAll(c *gin.Context) {

	name := c.Query("name")
	result, err := new(models.Category).CategoryFindAll(name)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "Database query error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "search successful",
	})

}

func CategoryCreate(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err := category.Create()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "Database creation error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "Created successfully",
	})
}

func CategoryUpdateField(c *gin.Context) {

	id := c.Query("id")
	field := c.Query("field")
	fieldContent := c.Query("fieldContent")

	err := new(models.Category).UpdateFieldContent(id, field, fieldContent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "update completed",
	})
}

func CategoryDestroy(c *gin.Context) {

	id := c.Query("id")

	err := new(models.Category).Destroy(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   id,
			"msg":    err.Error(),
		})
		return
	}
	claims, _ := serviceUser.GetJwtClaimsByContext(c)
	msg := models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "”deleted the category",
	}
	models.Broadcast <- msg

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
	})

}
