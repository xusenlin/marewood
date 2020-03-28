package controller

import (
	"MareWood/sql"
	"MareWood/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CategoryFindAll(c *gin.Context) {

	var result []models.Category

	if sql.DB.Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
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
	
	if sql.DB.Create(&category).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   sql.DB.Error.Error(),
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

	id := c.Query("id")

	var jobCount int

	if sql.DB.Model(&models.Job{}).Where("category_id = ?", id).Count(&jobCount).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   id,
			"msg":    sql.DB.Error.Error(),
		})
	}

	if jobCount > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   jobCount,
			"msg":    "无法删除，还有" + strconv.Itoa(jobCount) + "个任务在使用此分类",
		})
		return
	}

	if sql.DB.Where("id = ?", id ).Delete(&models.Category{}).Error != nil {
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
