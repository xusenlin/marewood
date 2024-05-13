package repository

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	ginPagination "github.com/xusenlin/gin-pagination"
	"io"
	"io/ioutil"
	"log/slog"
	"marewood/conf"
	"marewood/internal/command"
	"marewood/internal/common"
	"marewood/internal/context"
	"marewood/internal/db"
	"marewood/internal/event"
	"marewood/internal/log"
	"path/filepath"
	"strconv"
)

func FindAll(c *gin.Context) {

	ctx := context.New(c)

	var repositories []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Url  string `json:"url"`
		Tip  string `json:"tip"`
	}

	err := db.Conn.Model(Repository{}).Select("id", "name", "url").Find(&repositories).Error

	if err != nil {
		ctx.SendErr(err)
		return
	}
	for idx, r := range repositories {
		name, err := getRepositoryNameByUrl(r.Url)
		if err != nil {
			repositories[idx].Tip = "git url err"
		} else {
			repositories[idx].Tip = name
		}
	}
	ctx.SendOk("success", repositories)
}

func Find(c *gin.Context) {

	ctx := context.New(c)
	model := new(Repository)

	pagination := ginPagination.New[*Repository](model, c)

	pagination.Eq("id").Like("name")

	err := pagination.Query()

	if err != nil {
		ctx.SendErr(err)
		return
	}

	ctx.SendOk("pagination success", pagination)
}

func Edit(c *gin.Context) {

	var repository Repository
	ctx := context.New(c)
	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = c.ShouldBindJSON(&repository)
	if err != nil {
		ctx.SendErr(err)
		return
	}

	err = repository.Edit()
	if err != nil {
		ctx.SendErr(err)
		return
	}
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
	go repository.Clone(claims)
	ctx.SendOk("edit success", repository.ID)
}

func EventSource(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	claims, err := common.JwtParseToken(c.Query("token"))

	if err != nil {
		c.Stream(func(w io.Writer) bool {
			c.SSEvent("error", "token error")
			return false
		})
		return
	}

	event.RepoSource.Subscribe(claims.ID)
	defer event.RepoSource.CancelSubscribe(claims.ID)
	log.Slog.Info("[RepoEventSource]:subscribe", slog.Any("userId", claims.ID))
	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-event.RepoSource.ReceiveMsg(claims.ID):
			c.SSEvent(msg.Type, msg.Data)
			log.Slog.Info("[RepoEventSource]:send", slog.Any("userId", claims.ID), slog.String("type", msg.Type), slog.String("msg", msg.Data.Msg))
			return true
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}

func Destroy(c *gin.Context) {
	ctx := context.New(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil || id == 0 {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}

	var repo Repository
	repo.ID = uint(id)

	claims, err := ctx.GetClaims()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = repo.Destroy(claims)

	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("successfully deleted", id)
}

func GitPull(c *gin.Context) {
	ctx := context.New(c)
	id := c.Param("id")

	if id == "" {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}
	dst := filepath.Join(conf.RepositoryDir, id)

	out, err := command.GitPull(dst)
	if err != nil {
		ctx.SendErr(errors.New(out))
		return
	}
	ctx.SendOk("", out)
}

func GitCheckoutDot(c *gin.Context) {
	ctx := context.New(c)
	id := c.Param("id")

	if id == "" {
		ctx.SendErr(errors.New("id is not allowed to be empty"))
		return
	}
	dst := filepath.Join(conf.RepositoryDir, id)

	out, err := command.GitCheckoutDot(dst)
	if err != nil {
		ctx.SendErr(errors.New(out))
		return
	}
	ctx.SendOk("", out)
}

func GetPackageJson(c *gin.Context) {

	ctx := context.New(c)
	repositoryId := c.Param("repositoryId")
	if repositoryId == "" {
		ctx.SendErr(errors.New("repositoryId is not allowed to be empty"))
		return
	}
	dst := filepath.Join(conf.RepositoryDir, repositoryId)
	out, err := command.GitPull(dst)

	if err != nil {
		ctx.SendErr(errors.New(out))
		return
	}

	var packageJson PackageJson
	packageJsonFile := filepath.Join(dst, "package.json")
	content, err := ioutil.ReadFile(packageJsonFile)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	err = json.Unmarshal(content, &packageJson)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("success", packageJson)
}
func Branch(c *gin.Context) {

	ctx := context.New(c)
	repositoryId := c.Param("repositoryId")

	if repositoryId == "" {
		ctx.SendErr(errors.New("repositoryId is not allowed to be empty"))
		return
	}
	dst := filepath.Join(conf.RepositoryDir, repositoryId)

	out, err := command.GitPull(dst)

	if err != nil {
		ctx.SendErr(errors.New(out))
		return
	}
	branches, err := command.GitBranchList(dst)
	if err != nil {
		ctx.SendErr(err)
		return
	}
	ctx.SendOk("success", branches)
}
