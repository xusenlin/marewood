package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindAll(c *gin.Context) {

	var result []models.Category

	if database.DB.Find(&result).Error != nil {
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

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	category.JobQuantity = 0
	
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

	if true {//这里后面要根据此分类是否被使用来确认分类是否能被删除
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   id,
			"msg":    "无法删除，还有任务在使用此分类",
		})
		return
	}


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
