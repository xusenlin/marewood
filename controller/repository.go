package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"FEDeployService/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

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

func RepositoryCreate(c *gin.Context) {

	var repository models.Repository

	if err := c.BindJSON(&repository); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "出错",
		})
		return
	}

	repository.WebHookSecret = service.RandSeq(models.RepoWebHookSecretRandSeqLen)
	repository.Status = models.RepoStatusProcessing
	repository.DependStatus = models.RepoDependStatusProcessing

	if database.DB.Create(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error.Error(),
			"msg":    "数据库创建出错",
		})
		return
	}

	go func() {

		out, err := service.GitClone(repository.Url, repository.UserName, repository.Password)
		if err != nil {
			database.DB.Model(&repository).
				Where("id = ?", repository.ID).
				Update("status", models.RepoStatusFail).
				Update("terminal_info", out)
			return
		}

		database.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("status", models.RepoStatusSuccess).
			Update("terminal_info", out)



	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "后台正在执行仓库克隆",
		"msg":    "创建成功",
	})
}

//不同时删除仓库目录，因为任务可能还在使用
func RepositoryDestroy(c *gin.Context) {
	var repository models.Repository

	id := c.Query("id")

	if database.DB.Where("id = ?", id).Delete(&repository).Error != nil {
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

//webHook 更新
func RepositoryUpdate(c *gin.Context) {
	repositoryId := c.Query("id")

	if "push" != c.GetHeader("x-github-event") {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "webHook不是push事件",
		})
		return
	}

	var signature string

	signature = c.GetHeader("HTTP_X_GITLAB_TOKEN") //GitLab
	if "" == signature{
		signature = c.GetHeader("X-Hub-Signature") //Github
	}
	if "" == signature{
		signature = c.GetHeader("password") //Gitee
	}

	bodyContent,err := c.GetRawData()

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   err.Error(),
			"msg":    "bodyContent获取出错",
		})
		return
	}

	var repository models.Repository

	if database.DB.First(&repository, repositoryId).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   database.DB.Error,
			"msg":    "通过id找不到仓库",
		})
		return
	}

	if !service.VerificationWebHookSecret(repository.WebHookSecret,signature,bodyContent){
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "密钥验证不通过",
		})
		return
	}

	go func() {
		service.GitPullAndSaveRecord(repository.Url,repository.ID)
	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "git pull 已经执行",
	})
}
