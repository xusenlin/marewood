package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindAll(c *gin.Context) {

	var result []models.Category

	if database.DB.Order("created_at desc").Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
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

	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "出错",
		})
		return
	}

	if database.DB.Create(&category).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
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

func CategoryDestroy(c *gin.Context) {

	var category models.Category
	id := c.Query("id")

	//删除之前请确保此分类没有被使用
	if database.DB.Where("id = ?", id).Delete(&category).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库删除出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}
