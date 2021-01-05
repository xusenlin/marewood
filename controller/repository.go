package controller

import (
	"MareWood/models"
	"MareWood/service/serviceRepository"
	"MareWood/service/serviceUser"
	"MareWood/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
		"msg":    "search successful",
	})

}

func RepositoryFind(c *gin.Context) {

	name := c.Query("name")
	pageNum, err := strconv.Atoi(c.Query("pageNum"))

	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 8
	}
	var result models.RepositoryPageResult

	if strings.HasPrefix(name, "id:") {
		result, err = models.FindRepository("", pageNum, pageSize, map[string]interface{}{"id": strings.TrimLeft(name, "id:")})
	} else {
		result, err = models.FindRepository(name, pageNum, pageSize, map[string]interface{}{})
	}

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
		"msg":    "search successful",
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
		Message:         "“" + claims.Username + "”created the repository“" + repository.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "Created successfully, repository clone is being executed in the background",
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
			"msg":    "Unable to delete, there are " + strconv.Itoa(jobCount) + " tasks using this repository",
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
			"msg":    "Unable to delete the repository being cloned, please try again later",
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
		Message:         "“" + claims.Username + "”delete repository“" + repository.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "successfully deleted",
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
		"msg":    "execution succeed",
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
		"msg":    "execution succeed",
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
		"data":   "The repository dependency has been successfully deleted and will be reinstalled when performing related tasks",
		"msg":    "execution succeed",
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
		"msg":    "execution succeed",
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
		"msg":    "Branch has been cut",
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
		"msg":    "update completed",
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
			"msg":    "A task is being packaged, please complete the package and try again",
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
		Message:         "“" + claims.Username + "” rebuild the repository“" + repo.Name + "”",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "rebuilding the repository in the background...",
	})
}
