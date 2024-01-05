package repository

import (
	"errors"
	"gorm.io/gorm"
	"marewood/conf"
	"marewood/internal/command"
	"marewood/internal/common"
	"marewood/internal/db"
	"marewood/internal/event"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Repository struct {
	common.Model
	Name          string `binding:"required,min=2,max=40" json:"name"`
	Description   string `gorm:"type:varchar(1000)"  json:"description"`
	Url           string `binding:"required,url" json:"url"`     //仓库的地址
	UserName      string `json:"userName"`                       //仓库是私有的话需要填写
	Password      string `json:"password"`                       //仓库是私有的话需要填写
	Status        int    `json:"status"`                         //此仓库是否已经将代码克隆过来
	TaskStatus    int    `json:"taskStatus"`                     //有任务正在打包，其他任务无法执行
	DependTools   string `binding:"required" json:"dependTools"` //依赖工具选择
	DefaultBranch string `binding:"required,min=1,max=40" json:"defaultBranch"`
	TerminalInfo  string `gorm:"type:varchar(1000)" json:"terminalInfo"`
}

type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
}

func (r *Repository) Edit() error {

	query := db.Conn.Where("url = ?", r.Url)
	if r.ID != 0 { //edit
		query = query.Where("id != ?", r.ID)
	}
	err := query.First(&Repository{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("repository url already exists")
	}

	r.Status = StatusProcessing
	r.TaskStatus = TaskStatusLeisured
	r.TerminalInfo = ""

	if r.ID != 0 {
		return db.Conn.Save(&r).Error
	} else {
		return db.Conn.Create(&r).Error
	}
}

func (r *Repository) Clone(claims *common.Claims) {
	repoDir := r.CodeDir()
	terminalInfo := ""
	_ = os.RemoveAll(repoDir)
	terminalInfo += "RemoveDir:" + repoDir
	out, err := command.GitClone(r.UserName, r.Password, r.DefaultBranch, r.Url, repoDir)
	terminalInfo += "\n" + out
	eventType := ""
	eventData := &event.RepoData{
		RepositoryId: r.ID,
		UserId:       claims.ID,
		UserName:     claims.Username,
	}

	if err == nil {
		r.Status = StatusSuccess
		eventType = event.RepoTypeCloneOk
		eventData.Msg = "clone success"
		terminalInfo += "\n🥳🥳🥳🥳 Clone successful  🌼🌼🌼"
	} else {
		r.Status = StatusFail
		eventType = event.RepoTypeCloneErr
		eventData.Msg = err.Error()
		terminalInfo += "\n" + err.Error()
	}
	r.TerminalInfo = terminalInfo
	db.Conn.Save(r)

	event.RepoSource.PublishMsg(eventType, eventData)
}

func (r *Repository) Destroy(c *common.Claims) error {
	if errors.Is(db.Conn.First(&r).Error, gorm.ErrRecordNotFound) {
		return errors.New("repository not found")
	}

	var taskCount int64
	if err := db.Conn.Table("tasks").Where("repository_id =?", r.ID).Count(&taskCount).Error; err != nil {
		return err
	}
	if taskCount != 0 {
		return errors.New("the repository has tasks, cannot be deleted")
	}

	if err := db.Conn.Delete(&r).Error; err != nil {
		return err
	}
	event.RepoSource.PublishMsgExcludeID(event.RepoTypeDestroyOk, &event.RepoData{
		RepositoryId: r.ID,
		UserId:       c.ID,
		UserName:     c.Username,
		Msg:          "destroy success",
	}, c.ID)

	return os.RemoveAll(r.CodeDir())
}

func getRepositoryNameByUrl(gitUrl string) (string, error) {

	if !strings.HasPrefix(gitUrl, "http") {
		return "", errors.New("the address must have a http prefix! ")
	}

	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")

	urlArr := strings.Split(noSuffixUrl, "/")

	return urlArr[len(urlArr)-1], nil
}

func (r *Repository) CodeDir() string {
	return filepath.Join(conf.RepositoryDir, strconv.Itoa(int(r.ID)))
}
