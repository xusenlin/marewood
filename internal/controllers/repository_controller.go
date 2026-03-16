package controllers

import (
	"io"
	"marewood/internal/models"
	"marewood/internal/pkg/context"
	pkgErrors "marewood/internal/pkg/errors"
	"marewood/internal/pkg/event"
	"marewood/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
)

// RepositoryController 仓库控制器
type RepositoryController struct {
	repoService services.RepositoryService
}

// NewRepositoryController 创建仓库控制器
func NewRepositoryController(repoService services.RepositoryService) *RepositoryController {
	return &RepositoryController{
		repoService: repoService,
	}
}

// FindAll 查询所有仓库（简略信息）
func (ctrl *RepositoryController) FindAll(c *gin.Context) {
	ctx := context.New(c)

	repos, err := ctrl.repoService.ListAll()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("success", repos)
}

// Find 查询仓库列表（分页）
func (ctrl *RepositoryController) Find(c *gin.Context) {
	ctx := context.New(c)

	model := new(models.Repository)
	pagination := ginPagination.New[*models.Repository](model, c)

	pagination.Eq("id").Like("name")

	err := pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("pagination success", pagination)
}

// Edit 创建或编辑仓库
func (ctrl *RepositoryController) Edit(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	var repository models.Repository
	err = c.ShouldBindJSON(&repository)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	err = ctrl.repoService.CreateOrUpdate(&repository)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	// 发布编辑成功事件
	event.RepoSource.PublishMsgExcludeID(
		event.RepoTypeEditOk,
		&event.RepoData{
			RepositoryId: repository.ID,
			UserId:       claims.ID,
			UserName:     claims.Username,
			Msg:          "edit success",
		},
		claims.ID,
	)

	// 异步克隆仓库
	go ctrl.repoService.Clone(repository.ID, claims)

	ctx.SendOk("edit success", repository.ID)
}

// Destroy 删除仓库
func (ctrl *RepositoryController) Destroy(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.SendErr(pkgErrors.NewValidationError("id is not allowed to be empty"))
		return
	}

	err = ctrl.repoService.DeleteRepository(uint(id), claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("successfully deleted", id)
}

// GitPull 执行 Git Pull
func (ctrl *RepositoryController) GitPull(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.SendErr(pkgErrors.NewValidationError("id is not allowed to be empty"))
		return
	}

	err = ctrl.repoService.GitPull(uint(id), claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("git pull success", id)
}

// GitCheckoutDot 执行 Git Checkout .
func (ctrl *RepositoryController) GitCheckoutDot(c *gin.Context) {
	ctx := context.New(c)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.SendErr(pkgErrors.NewValidationError("id is not allowed to be empty"))
		return
	}

	err = ctrl.repoService.GitCheckoutDot(uint(id), claims)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("git checkout . success", id)
}

// Branch 获取仓库分支列表
func (ctrl *RepositoryController) Branch(c *gin.Context) {
	ctx := context.New(c)

	repositoryId, err := strconv.ParseUint(c.Param("repositoryId"), 10, 32)
	if err != nil || repositoryId == 0 {
		ctx.SendErr(pkgErrors.NewValidationError("repositoryId is not allowed to be empty"))
		return
	}

	branches, err := ctrl.repoService.GetBranches(uint(repositoryId))
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("get branches success", branches)
}

// GetPackageJson 获取 package.json
func (ctrl *RepositoryController) GetPackageJson(c *gin.Context) {
	ctx := context.New(c)

	repositoryId, err := strconv.ParseUint(c.Param("repositoryId"), 10, 32)
	if err != nil || repositoryId == 0 {
		ctx.SendErr(pkgErrors.NewValidationError("repositoryId is not allowed to be empty"))
		return
	}

	packageJson, err := ctrl.repoService.GetPackageJson(uint(repositoryId))
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("get package.json success", packageJson)
}

// EventSource SSE 事件流
func (ctrl *RepositoryController) EventSource(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	claims, err := context.New(c).GetClaims()
	var userID uint
	if err != nil {
		userID = 0
	} else {
		userID = claims.ID
	}

	event.RepoSource.Subscribe(userID)

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-event.RepoSource.ReceiveMsg(userID):
			c.SSEvent(msg.Type, msg.Data)
			return true
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}
