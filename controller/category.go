package controller

import (
	"MareWood/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindAll(c *gin.Context) {

	result, err := new(models.Category).CategoryFindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "数据库查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
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
			"msg":    "数据库创建出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建成功",
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
		"msg":    "更新成功",
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

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}
