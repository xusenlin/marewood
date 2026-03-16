package controllers

import (
	"marewood/internal/models"
	"marewood/internal/pkg/context"
	"marewood/internal/pkg/errors"
	"marewood/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
)

// UserController 用户控制器
type UserController struct {
	userService services.UserService
}

// NewUserController 创建用户控制器
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register 用户注册
func (ctrl *UserController) Register(c *gin.Context) {
	ctx := context.New(c)

	var req struct {
		Username string `json:"username" binding:"required,min=2,max=8"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ctx.SendErr(err)
		return
	}

	user, err := ctrl.userService.Register(req.Username, req.Password)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("registration success", user.ID)
}

// Login 用户登录
func (ctrl *UserController) Login(c *gin.Context) {
	ctx := context.New(c)

	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ctx.SendErr(err)
		return
	}

	token, user, err := ctrl.userService.Login(req.Username, req.Password)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("login successful", struct {
		models.User
		Token string `json:"token"`
	}{*user, token})
}

// Find 查询用户列表（分页）
func (ctrl *UserController) Find(c *gin.Context) {
	ctx := context.New(c)

	model := new(models.User)
	pagination := ginPagination.New[*models.User](model, c)

	pagination.Like("username").Eq("id")

	err := pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	// 隐藏密码
	for idx := range pagination.List {
		pagination.List[idx].Password = "***"
	}

	ctx.SendOk("search successful", pagination)
}

// Edit 编辑用户
func (ctrl *UserController) Edit(c *gin.Context) {
	ctx := context.New(c)

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		ctx.SendErr(err)
		return
	}

	// 获取原始密码（如果为空则不修改密码）
	var req struct {
		Password string `json:"password"`
	}
	c.ShouldBindJSON(&req)

	err := ctrl.userService.UpdateUser(&user, req.Password)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("Successfully modified", user.ID)
}

// Destroy 删除用户
func (ctrl *UserController) Destroy(c *gin.Context) {
	ctx := context.New(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.SendErr(errors.NewValidationError("id is not allowed to be empty"))
		return
	}

	err = ctrl.userService.DeleteUser(uint(id))
	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("successfully deleted", id)
}
