package controller

import (
	"MareWood/models"
	"MareWood/service/serviceRepository"
	"MareWood/service/serviceUser"
	"MareWood/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

	isNormal := c.Query("isNormal")

	query := sql.DB.Order("created_at desc").Select("id,name")

	if isNormal == "1" {
		query = query.Where("status = ?", models.RepoStatusSuccess)
	}

	if query.Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})

}

func RepositoryFind(c *gin.Context) {

	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 8
	}

	result, err := models.FindRepository(pageNum, pageSize, map[string]interface{}{})

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
		"data":   result,
		"msg":    "查询成功",
	})
}

func RepositoryCreate(c *gin.Context) {

	var repository models.Repository

	claims, _ := serviceUser.GetJwtClaimsByContext(c)

	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err := repository.Create()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	go serviceRepository.CloneRepo(&repository, claims)

	models.Broadcast <- models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” 创建了仓库“" + repository.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建成功,后台正在执行仓库克隆",
	})
}

func RepositoryDestroy(c *gin.Context) {

	id := c.Query("id")

	var jobCount int

	if sql.DB.Model(&models.Job{}).Where("repository_id = ?", id).Count(&jobCount).Error != nil {
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
			"msg":    "无法删除，还有" + strconv.Itoa(jobCount) + "个任务在使用此仓库",
		})
		return
	}

	var repository models.Repository

	if sql.DB.First(&repository, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	if repository.Status == models.RepoStatusProcessing {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "无法删除正在克隆的仓库，请稍后再试",
		})
		return
	}

	err := serviceRepository.DeleteRepository(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	if sql.DB.Delete(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	claims, _ := serviceUser.GetJwtClaimsByContext(c)
	models.Broadcast <- models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” 删除了仓库“" + repository.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}

func RepositoryGitPull(c *gin.Context) {

	id := c.Query("id")
	out, err := serviceRepository.GitPull(id)

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
		"data":   out,
		"msg":    "执行成功",
	})
}

func RepositoryDiscardChange(c *gin.Context) {

	id := c.Query("id")
	out, err := serviceRepository.DiscardChange(id)

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
		"data":   out,
		"msg":    "执行成功",
	})
}

func RepositoryDeleteDepend(c *gin.Context) {

	id := c.Query("id")

	err := serviceRepository.DeleteDepend(id)

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
		"data":   "仓库依赖删除成功，相关任务执行的时候会重新安装",
		"msg":    "执行成功",
	})
}

func RepositoryBranch(c *gin.Context) {

	id := c.Query("id")

	branch, err := serviceRepository.GetBranch(id)

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
		"data":   branch,
		"msg":    "执行成功",
	})

}

func RepositoryPruneBranch(c *gin.Context) {

	id := c.Query("id")

	out, err := serviceRepository.PruneBranch(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    out,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "分支已经裁剪",
	})

}

func RepositoryScript(c *gin.Context) {
	id := c.Query("id")
	packageJson, err := serviceRepository.ParsePackageJson(id)
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
		"data":   packageJson.Scripts,
		"msg":    "",
	})
}

func RepositoryUpdateField(c *gin.Context) {

	id := c.Query("id")
	field := c.Query("field")
	fieldContent := c.Query("fieldContent")

	err := new(models.Repository).UpdateFieldContent(id, field, fieldContent)

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

func RepositoryReset(c *gin.Context) {

	id := c.Query("id")
	var repo models.Repository

	query := sql.DB.First(&repo, id)
	if query.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}
	if repo.JobStatus != models.RepoJobStatusLeisured {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "有任务正在打包，请打包完成再试",
		})
		return
	}
	//更新 status ，防止任务在这个时候打包
	if query.Update("status", models.RepoStatusProcessing).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	claims, _ := serviceUser.GetJwtClaimsByContext(c)

	err := serviceRepository.DeleteRepository(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	go serviceRepository.CloneRepo(&repo, claims)

	models.Broadcast <- models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” 重建了仓库“" + repo.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "后台重建仓库中...",
	})
}
