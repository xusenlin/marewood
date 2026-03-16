package services

import (
	"crypto/md5"
	"fmt"
	"marewood/internal/dao"
	"marewood/internal/models"
	"marewood/internal/pkg/errors"
	"marewood/internal/pkg/jwt"
)

// UserService 用户业务逻辑接口
type UserService interface {
	Register(username, password string) (*models.User, error)
	Login(username, password string) (string, *models.User, error)
	GetUser(id uint) (*models.User, error)
	UpdateUser(user *models.User, password string) error
	DeleteUser(id uint) error
	ListUsers(page, pageSize int, filters map[string]interface{}) ([]*models.User, int64, error)
}

type userService struct {
	userDAO dao.UserDAO
}

// NewUserService 创建用户服务
func NewUserService(userDAO dao.UserDAO) UserService {
	return &userService{
		userDAO: userDAO,
	}
}

func (s *userService) Register(username, password string) (*models.User, error) {
	// 检查用户名是否存在
	existing, err := s.userDAO.FindByUsername(username)
	if err == nil && existing != nil {
		return nil, errors.NewConflictError("username already exists")
	}

	// 创建用户
	user := &models.User{
		Username: username,
		Password: md5Hash(password),
		Status:   models.UserStatusEnabled,
	}

	// 设置角色（Admin 用户为超管）
	if username == "Admin" {
		user.Role = models.UserRoleSuperAdmin
	} else {
		user.Role = models.UserRoleReporter
	}

	if err := s.userDAO.Create(user); err != nil {
		return nil, errors.NewInternalError(err)
	}

	return user, nil
}

func (s *userService) Login(username, password string) (string, *models.User, error) {
	// 查找用户
	user, err := s.userDAO.FindByUsername(username)
	if err != nil {
		if _, ok := err.(*errors.AppError); ok {
			return "", nil, errors.NewValidationError("username or password is wrong")
		}
		return "", nil, errors.NewInternalError(err)
	}

	// 验证密码
	if user.Password != md5Hash(password) {
		return "", nil, errors.NewValidationError("username or password is wrong")
	}

	// 检查状态
	if user.Status == models.UserStatusDisabled {
		return "", nil, errors.NewValidationError("user is disabled")
	}

	// 生成 Token
	token, err := jwt.GenerateToken(&jwt.Claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	})
	if err != nil {
		return "", nil, errors.NewInternalError(err)
	}

	// 隐藏密码
	user.Password = "***"

	return token, user, nil
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.userDAO.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User, password string) error {
	// 查找旧用户
	oldUser, err := s.userDAO.FindByID(user.ID)
	if err != nil {
		return err
	}

	// 不能编辑超管
	if oldUser.Role == models.UserRoleSuperAdmin {
		return errors.NewValidationError("can not edit super administrator")
	}

	// 处理密码
	if password == "" {
		user.Password = oldUser.Password
	} else {
		user.Password = md5Hash(password)
	}

	return s.userDAO.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	// 查找用户
	user, err := s.userDAO.FindByID(id)
	if err != nil {
		return err
	}

	// 不能删除超管
	if user.Role == models.UserRoleSuperAdmin {
		return errors.NewValidationError("can not delete super administrator")
	}

	return s.userDAO.Delete(id)
}

func (s *userService) ListUsers(page, pageSize int, filters map[string]interface{}) ([]*models.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.userDAO.List(offset, pageSize, filters)
}

// md5Hash MD5 加密
func md5Hash(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(models.PasswordSalt+str)))
}
