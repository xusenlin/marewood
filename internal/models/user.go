package models

// User 用户模型
type User struct {
	Model
	Username    string `gorm:"unique" binding:"required,min=2,max=8" json:"username"`
	Password    string `json:"password"`
	Role        int    `json:"role"`
	Status      int    `gorm:"default:1" json:"status"`
	Avatar      string `json:"avatar"`
	Description string `gorm:"type:varchar(200)" json:"description"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

const ( //用户角色
	UserRoleReporter   = 10
	UserRoleDeveloper  = 6
	UserRoleAdmin      = 3
	UserRoleSuperAdmin = 1
)

// 用户状态常量
const (
	UserStatusEnabled  = 1
	UserStatusDisabled = 0
)

// PasswordSalt 密码盐值
const PasswordSalt = "marewood"
