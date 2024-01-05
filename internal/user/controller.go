package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
	"marewood/internal/common"
	"marewood/internal/context"
	"strconv"
)

func Register(c *gin.Context) {
	var newUser User
	ctx := context.New(c)
	if err := c.ShouldBindJSON(&newUser); err != nil {
		ctx.SendErr(err)
		return
	}

	if err := newUser.Register(); err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("registration success", newUser.ID)
}

func Login(c *gin.Context) {

	var user User
	ctx := context.New(c)

	if err := c.ShouldBindJSON(&user); err != nil {
		ctx.SendErr(err)
		return
	}

	if err := user.Login(); err != nil {
		ctx.SendErr(err)
		return
	}

	token, err := common.JwtGenToken(&common.Claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	})

	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("login successful", struct {
		User
		Token string `json:"token"`
	}{user, token})
}

func Find(c *gin.Context) {

	model := new(User)
	ctx := context.New(c)

	pagination := ginPagination.New[*User](model, c)

	pagination.Like("username").Eq("id")

	err := pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	for idx := range pagination.List {
		pagination.List[idx].Password = "***"
	}

	ctx.SendOk("search successful", pagination)
}

func Destroy(c *gin.Context) {
	ctx := context.New(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}

	var user User
	user.ID = uint(id)

	err = user.Destroy()

	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("successfully deleted", id)
}

func Edit(c *gin.Context) {

	var user User
	ctx := context.New(c)

	if err := c.ShouldBindJSON(&user); err != nil {
		ctx.SendErr(err)
		return
	}

	err := user.Edit()

	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("Successfully modified", user.ID)
}
