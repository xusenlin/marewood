package models

import (
	"marewood/conf"
	"path/filepath"
	"strconv"
)

// Repository 仓库模型
type Repository struct {
	Model
	Name          string `binding:"required,min=2,max=40" json:"name"`
	Description   string `gorm:"type:varchar(1000)" json:"description"`
	Url           string `binding:"required,url" json:"url"`     // 仓库的地址
	UserName      string `json:"userName"`                       // 仓库是私有的话需要填写
	Password      string `json:"password"`                       // 仓库是私有的话需要填写
	Status        int    `json:"status"`                         // 此仓库是否已经将代码克隆过来
	TaskStatus    int    `json:"taskStatus"`                     // 有任务正在打包，其他任务无法执行
	DependTools   string `binding:"required" json:"dependTools"` // 依赖工具选择
	DefaultBranch string `binding:"required,min=1,max=40" json:"defaultBranch"`
	TerminalInfo  string `gorm:"type:varchar(1000)" json:"terminalInfo"`
}

// TableName 指定表名
func (Repository) TableName() string {
	return "repositories"
}

// CodeDir 获取代码目录
func (r *Repository) CodeDir() string {
	return filepath.Join(conf.RepositoryDir, strconv.Itoa(int(r.ID)))
}

// 仓库状态常量
const (
	RepositoryStatusProcessing = 0 // 正在克隆
	RepositoryStatusSuccess    = 1 // 克隆成功
	RepositoryStatusFail       = 2 // 克隆失败
)

// 任务状态常量
const (
	RepositoryTaskStatusLeisured = 0 // 空闲
	RepositoryTaskStatusBusy     = 1 // 忙碌
)

// PackageJson package.json 结构
type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
}
