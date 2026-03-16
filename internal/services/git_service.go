package services

import (
	"marewood/internal/pkg/command"
)

// GitService Git 操作服务接口
type GitService interface {
	Clone(username, password, branch, url, dir string) (string, error)
	Pull(dir string) (string, error)
	GetBranches(dir string) ([]string, error)
	CheckoutDot(dir string) (string, error)
	GetCommitHash(dir string) (string, error)
}

type gitService struct{}

// NewGitService 创建 Git 服务
func NewGitService() GitService {
	return &gitService{}
}

func (s *gitService) Clone(username, password, branch, url, dir string) (string, error) {
	return command.GitClone(username, password, branch, url, dir)
}

func (s *gitService) Pull(dir string) (string, error) {
	return command.GitPull(dir)
}

func (s *gitService) GetBranches(dir string) ([]string, error) {
	return command.GitBranchList(dir)
}

func (s *gitService) CheckoutDot(dir string) (string, error) {
	return command.GitCheckoutDot(dir)
}

func (s *gitService) GetCommitHash(dir string) (string, error) {
	return command.GitCommitHash(dir)
}
