package models

import (
	"marewood/conf"
	"path/filepath"
	"strconv"
)

// TaskStatus 任务状态类型
type TaskStatus uint

const (
	TaskStatusDefault    TaskStatus = 0 // 默认
	TaskStatusProcessing TaskStatus = 1 // 执行中
	TaskStatusFailed     TaskStatus = 2 // 失败
	TaskStatusSuccess    TaskStatus = 3 // 成功
)

type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Task 任务模型
type Task struct {
	Model
	Name         string     `binding:"required,min=2,max=20" json:"name"`
	Description  string     `gorm:"type:varchar(1000)" json:"description"`
	Status       TaskStatus `gorm:"default:0" json:"status"`
	Branch       string     `binding:"required,min=1,max=20" json:"branch"`
	Alias        string     `binding:"required,min=2,max=100" json:"alias"` // 打包成功之后的url访问目录
	RunTotal     uint       `gorm:"default:0" json:"runTotal"`
	RepositoryId uint       `gorm:"index" binding:"required" json:"repositoryId"`
	BuildDir     string     `gorm:"default:'dist'" binding:"required,min=2,max=20" json:"buildDir"` // 打包的目录,默认是dist
	BuildCommand string     `binding:"required,min=2,max=30" json:"buildCommand"`                   // 打包命令
	CreatedId    uint       `json:"createdId"`                                                      // 创建的用户ID
	Private      bool       `json:"private"`                                                        // 私密任务，仅自己可见
	Tag          string     `json:"tag"`
	CommitHash   string     `json:"commitHash"` // 打包时的commit hash
	TerminalInfo string     `gorm:"type:varchar(1000)" json:"terminalInfo"`
}

// TableName 指定表名
func (Task) TableName() string {
	return "tasks"
}

// WebDir 获取 Web 目录
func (t *Task) WebDir() string {
	return filepath.Join(conf.WebRootDir, t.Alias)
}

// RepositoryCodeDir 获取仓库代码目录
func (t *Task) RepositoryCodeDir() string {
	return filepath.Join(conf.RepositoryDir, strconv.Itoa(int(t.RepositoryId)))
}
