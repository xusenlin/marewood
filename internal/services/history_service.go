package services

import (
	"fmt"
	"marewood/conf"
	"marewood/internal/dao"
	"marewood/internal/models"
	"os"
	"path/filepath"
	"strconv"
)

type HistoryService interface {
	FindByTaskID(taskId uint) ([]models.History, error)
	Restore(id uint) error
}

type historyService struct {
	historyDao dao.HistoryDao
	taskDao    dao.TaskDAO
}

func NewHistoryService(historyDao dao.HistoryDao, taskDao dao.TaskDAO) HistoryService {
	return &historyService{historyDao: historyDao, taskDao: taskDao}
}

func (s *historyService) FindByTaskID(taskId uint) ([]models.History, error) {
	return s.historyDao.FindByTaskID(taskId)
}

func (s *historyService) Restore(id uint) error {
	history, err := s.historyDao.FindByID(id)
	if err != nil {
		return err
	}

	task, err := s.taskDao.FindByID(history.TaskID)
	if err != nil {
		return err
	}

	historyDir := filepath.Join(conf.HistoryDir, strconv.Itoa(int(task.ID)), history.GitHash)
	if _, err := os.Stat(historyDir); os.IsNotExist(err) {
		return fmt.Errorf("history version not found: %s", history.GitHash)
	}

	// 恢复操作
	webDir := task.WebDir()
	if err := os.RemoveAll(webDir); err != nil {
		return err
	}

	if err := s.copy(historyDir, webDir); err != nil {
		return err
	}

	task.CommitHash = history.GitHash

	return s.taskDao.Update(task)
}

func (s *historyService) copy(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(dstPath, data, info.Mode())
	})
}
