package controller

import (
	"FrontEndAutomatedDeployment/database"
	"FrontEndAutomatedDeployment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindAll(c *gin.Context)  {

	var result []models.Category

	database.DB.Order("created_at desc").Find(&result)

	if database.DB.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": result,
		"msg":"查询成功",
	})

}

func CategoryCreate(c *gin.Context)  {

	var category models.Category

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": err.Error(),
			"msg":"出错",
		})
		return
	}

	database.DB.Create(&category)

	if database.DB.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库创建出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":"",
		"msg":"创建成功",
	})
}

func CategoryDestroy(c *gin.Context)  {
	var category models.Category

	id := c.Query("id")
	//删除之前请确保此分类没有被使用
	database.DB.Where("id = ?", id ).Delete(&category)

	if database.DB.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data": database.DB.Error.Error(),
			"msg":"数据库删除出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":id,
		"msg":"删除成功",
	})

}