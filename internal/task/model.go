package task

import (
	"archive/tar"
	"bytes"
	"errors"
	"gorm.io/gorm"
	"io"
	"marewood/conf"
	"marewood/internal/command"
	"marewood/internal/common"
	"marewood/internal/db"
	"marewood/internal/event"
	"os"
	"path/filepath"
	"strconv"
)

type Task struct {
	common.Model
	Name         string     `binding:"required,min=2,max=20" json:"name"`
	Description  string     `gorm:"type:varchar(1000)"  json:"description"`
	Status       taskStatus `gorm:"default:0" json:"status"`
	Branch       string     `binding:"required,min=1,max=20"  json:"branch"`
	Alias        string     `binding:"required,min=2,max=100" json:"alias"` //打包成功之后的url访问目录`
	RunTotal     uint       `gorm:"default:0" json:"runTotal"`
	RepositoryId uint       `gorm:"index" binding:"required" json:"repositoryId"`
	BuildDir     string     `gorm:"default:'dist'" binding:"required,min=2,max=20" json:"buildDir"` //打包的目录,默认是dist
	BuildCommand string     `binding:"required,min=2,max=30" json:"buildCommand"`                   //打包命令，npm run build 可以读取package.json供选择
	CreatedId    uint       `json:"createdId"`                                                      //创建的用户ID
	Private      bool       `json:"private"`                                                        // 私密任务，仅自己可见
	Tag          string     `json:"tag"`
	CommitHash   string     `json:"commitHash"` //打包时的commit hash
	TerminalInfo string     `gorm:"type:varchar(1000)" json:"terminalInfo"`
}

func (t *Task) Destroy(c *common.Claims) error {
	if errors.Is(db.Conn.First(&t).Error, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}
	err := db.Conn.Delete(&t).Error
	if err != nil {
		return err
	}
	event.TaskSource.PublishMsgExcludeID(event.TaskTypeDestroyOk, &event.TaskData{
		TaskId:   t.ID,
		UserId:   c.ID,
		UserName: c.Username,
		Msg:      "destroy success",
	}, c.ID)
	return nil
}

func (t *Task) Edit(u *common.Claims) error {

	query := db.Conn.Where("alias = ?", t.Alias).Where("id <> ?", t.ID)
	err := query.First(&Task{}).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("alias url already exists")
	}

	t.Status = StatusDefault
	t.CreatedId = u.ID
	t.CommitHash = "None" //Branch change, commit hash will be None

	if t.ID != 0 {
		return db.Conn.Save(&t).Error
	} else {
		return db.Conn.Create(&t).Error
	}
}

func (t *Task) Tar() ([]byte, error) {

	dst := filepath.Join(conf.WebRootDir, t.Alias)
	folder, err := os.Open(dst)
	if err != nil {
		return nil, err
	}
	defer folder.Close()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// 遍历源目录下的文件和子目录
	err = filepath.Walk(dst, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 为当前文件或目录创建一个tar头
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// 更新tar头中的文件名
		rel, err := filepath.Rel(dst, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(rel)

		// 写入tar头
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// 如果当前文件是一个普通文件，则将内容写入tar文件
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return buf.Bytes(), err
	}

	// 关闭tar文件
	if err := tw.Close(); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}

func (t *Task) CheckBranch() error {

	dst := filepath.Join(conf.RepositoryDir, strconv.Itoa(int(t.RepositoryId)))
	branches, err := command.GitBranchList(dst)

	if err != nil {
		return err
	}
	for _, v := range branches {
		if v == t.Branch {
			return nil
		}
	}
	return errors.New("the branch does not exist")
}
