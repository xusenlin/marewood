# MareWood 架构文档

## 📐 架构概览

MareWood 采用经典的三层架构模式，结合依赖注入和接口驱动设计，确保代码的高内聚低耦合。

```
┌─────────────────────────────────────────────────────────┐
│                      HTTP Request                        │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│                   Middlewares 中间件                      │
│  • JWT 认证 (JWTAuth)                                    │
│  • 权限验证 (RoleReporter/Developer/SuperAdmin)          │
│  • 日志记录 (Logger/AuthLogger)                          │
│  • CORS 跨域                                             │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│              Controllers 控制器层                         │
│  • UserController - 用户 HTTP 处理                       │
│  • RepositoryController - 仓库 HTTP 处理                 │
│  • TaskController - 任务 HTTP 处理                       │
│  • HistoryController - 历史版本 HTTP 处理                 │
│                                                          │
│  职责：                                                   │
│  - 参数绑定和验证                                         │
│  - 调用 Service 执行业务                                  │
│  - 返回统一响应格式                                       │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│               Services 业务逻辑层                         │
│  • UserService - 用户注册、登录、权限管理                 │
│  • RepositoryService - 仓库克隆、Git 操作                │
│  • TaskService - 任务执行、打包、归档                     │
│  • HistoryService - 历史版本管理、版本恢复                │
│  • GitService - Git 命令封装                             │
│                                                          │
│  职责：                                                   │
│  - 核心业务规则                                          │
│  - 调用 DAO 进行数据操作                                  │
│  - 调用基础设施服务                                       │
│  - 发布事件通知                                          │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│                 DAO 数据访问层                            │
│  • UserDAO - User CRUD                                   │
│  • RepositoryDAO - Repository CRUD                       │
│  • TaskDAO - Task CRUD                                   │
│  • HistoryDAO - History CRUD                             │
│                                                          │
│  职责：                                                   │
│  - 数据库 CRUD 操作                                       │
│  - 查询条件封装                                          │
│  - 分页处理                                              │
│  - 事务管理                                              │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│                    Database 数据库                        │
│                   (SQLite + GORM)                        │
└─────────────────────────────────────────────────────────┘

        ┌───────────────────────────────────────┐
        │     Infrastructure (pkg/) 基础设施     │
        │  • database - 数据库连接管理           │
        │  • event - SSE 事件推送                │
        │  • command - Shell 命令执行            │
        │  • jwt - JWT 认证                     │
        │  • logger - 日志记录                   │
        │  • context - Gin 上下文工具            │
        │  • errors - 统一错误定义               │
        └───────────────────────────────────────┘
                 ↑            ↑
                 │            │
           Controllers    Services
```

---

## 🔄 数据流向

### 1. 用户登录流程

```
HTTP POST /v1/login
    ↓
[Middlewares] Logger
    ↓
[UserController] Login
    ├─ 参数绑定：username, password
    ├─ 调用 UserService.Login()
    │     ├─ UserDAO.FindByUsername()
    │     │     └─ SELECT * FROM users WHERE username = ?
    │     ├─ 验证密码（MD5）
    │     ├─ 检查用户状态
    │     └─ jwt.GenerateToken()
    └─ 返回 { user, token }
```

### 2. 仓库克隆流程

```
HTTP PUT /v1/repository
    ↓
[Middlewares] JWTAuth → AuthLogger → RoleDeveloper
    ↓
[RepositoryController] Edit
    ├─ 参数绑定：Repository
    ├─ 获取用户信息 (GetClaims)
    ├─ 调用 RepositoryService.CreateOrUpdate()
    │     ├─ RepositoryDAO.FindByURL() - 检查 URL 是否存在
    │     ├─ RepositoryDAO.Create() - 创建记录
    │     └─ 发布 event.RepoTypeEditOk
    ├─ 异步调用 RepositoryService.Clone()
    │     ├─ GitService.Clone() - 执行 git clone
    │     ├─ RepositoryDAO.Update() - 更新状态
    │     └─ 发布 event.RepoTypeCloneOk/CloneErr
    └─ 返回成功响应
```

### 3. 任务运行流程

```
HTTP GET /v1/task/run?id=1
    ↓
[Middlewares] JWTAuth → AuthLogger → RoleDeveloper
    ↓
[TaskController] Run
    ├─ 参数验证：taskId
    ├─ 获取用户信息 (GetClaims)
    ├─ 调用 TaskService.RunTask()
    │     ├─ TaskDAO.FindByID() - 获取任务
    │     ├─ 权限检查（私有任务）
    │     ├─ RepositoryDAO.FindByID() - 获取仓库
    │     ├─ 状态检查（仓库是否正常、是否空闲）
    │     ├─ 更新状态（任务Processing、仓库Busy）
    │     ├─ 发布 event.TaskTypeRunOk
    │     └─ 异步执行 executeTask()
    │           ├─ GitService.Pull()
    │           ├─ 检查分支
    │           ├─ GitService.Checkout()
    │           ├─ 安装依赖 (npm/yarn/pnpm)
    │           ├─ 执行构建 (npm run build)
    │           ├─ 移动文件到 web 目录
    │           ├─ GitService.GetCommitHash()
    │           ├─ 创建或更新历史版本记录
    │           │     ├─ HistoryDAO.FindByTaskIDAndGitHash() - 检查是否已存在
    │           │     ├─ 存在则更新用户信息
    │           │     └─ 不存在则创建新记录
    │           ├─ 复制构建文件到历史目录
    │           ├─ 更新任务状态 (Success/Failed)
    │           └─ 发布 event.TaskTypeBuildOk/BuildFail
    └─ 返回 "Compiling in the background..."
```

### 4. 历史版本恢复流程

```
HTTP POST /v1/history/restore/:id
    ↓
[Middlewares] JWTAuth → AuthLogger → RoleDeveloper
    ↓
[HistoryController] Restore
    ├─ 参数验证：historyId
    ├─ 调用 HistoryService.Restore()
    │     ├─ HistoryDAO.FindByID() - 获取历史版本记录
    │     ├─ TaskDAO.FindByID() - 获取关联任务
    │     ├─ 检查历史版本目录是否存在
    │     ├─ 删除当前 web 目录
    │     ├─ 复制历史版本文件到 web 目录
    │     └─ TaskDAO.Update() - 更新任务的 CommitHash
    └─ 返回成功响应
```

### 5. 查询历史版本流程

```
HTTP GET /v1/histories?taskId=1
    ↓
[Middlewares] JWTAuth → AuthLogger → RoleReporter
    ↓
[HistoryController] FindByTaskID
    ├─ 参数验证：taskId
    ├─ 调用 HistoryService.FindByTaskID()
    │     └─ HistoryDAO.FindByTaskID() - 查询最近 10 个版本
    └─ 返回历史版本列表
```

---

## 🏗️ 依赖关系图

### 模块依赖

```
Controllers
    ↓ 依赖
Services
    ↓ 依赖
DAOs & PKG
    ↓ 依赖
Models & Database
```

**严格规则：**
- Controllers **只能** 调用 Services
- Services **可以** 调用 DAOs 和 PKG
- DAOs **只能** 调用 Database 和 Models
- Models **不依赖** 任何其他层
- PKG **独立**，可被任何层调用

### 接口依赖注入

```go
// 在 routes/v1.go 中统一初始化

// 1. 获取数据库连接
db := database.GetDB()

// 2. 创建 DAO 实例（注入数据库连接）
userDAO := dao.NewUserDAO(db)
repoDAO := dao.NewRepositoryDAO(db)
taskDAO := dao.NewTaskDAO(db)
historyDAO := dao.NewHistoryDao(db)

// 3. 创建 Service 实例（注入 DAO 和其他 Service）
gitService := services.NewGitService()
userService := services.NewUserService(userDAO)
repoService := services.NewRepositoryService(repoDAO, gitService)
taskService := services.NewTaskService(taskDAO, repoDAO, historyDAO, gitService)
historyService := services.NewHistoryService(historyDAO, taskDAO)

// 4. 创建 Controller 实例（注入 Service）
userCtrl := controllers.NewUserController(userService)
repoCtrl := controllers.NewRepositoryController(repoService)
taskCtrl := controllers.NewTaskController(taskService)
historyCtrl := controllers.NewHistoryController(historyService)

// 5. 注册路由
v1.POST("/login", userCtrl.Login)
```

---

## 📦 模块职责详解

### Models 模块

**文件结构：**
```
models/
├── common.go       # Model 基类（ID、CreatedAt、UpdatedAt、DeletedAt）
├── user.go         # User 模型 + 常量（角色、状态）
├── repository.go   # Repository 模型 + 常量（状态）
├── task.go         # Task 模型 + 常量（状态）
└── history.go      # History 模型（版本记录）
```

**职责：**
- 定义数据结构
- GORM 标签（表名、字段类型、索引）
- JSON 标签（序列化）
- 验证标签（binding）
- 业务常量（状态码、角色等）

**不包含：**
- ❌ 业务逻辑
- ❌ 数据库操作
- ❌ 外部依赖

---

### DAO 模块

**文件结构：**
```
dao/
├── user_dao.go         # UserDAO 接口 + 实现
├── repository_dao.go   # RepositoryDAO 接口 + 实现
├── task_dao.go         # TaskDAO 接口 + 实现
└── history_dao.go      # HistoryDAO 接口 + 实现
```

**标准接口方法：**
```go
type UserDAO interface {
    Create(user *models.User) error
    FindByID(id uint) (*models.User, error)
    FindByUsername(username string) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
    List(offset, limit int, filters map[string]interface{}) ([]*models.User, int64, error)
}

type HistoryDAO interface {
    Create(history *models.History) error
    FindByID(id uint) (*models.History, error)
    FindByTaskID(taskId uint) ([]models.History, error)
    FindByTaskIDAndGitHash(taskId uint, gitHash string) (*models.History, error)
    Update(history *models.History) error
}
```

**职责：**
- 数据库 CRUD 操作
- 查询条件封装
- 分页逻辑
- 错误转换（GORM 错误 → AppError）

**不包含：**
- ❌ 业务验证
- ❌ 事件发布
- ❌ 外部服务调用

---

### Services 模块

**文件结构：**
```
services/
├── user_service.go       # 用户业务逻辑
├── repository_service.go # 仓库业务逻辑
├── task_service.go       # 任务业务逻辑
├── history_service.go    # 历史版本业务逻辑
└── git_service.go        # Git 操作封装
```

**标准接口方法：**
```go
type UserService interface {
    Register(username, password string) (*models.User, error)
    Login(username, password string) (string, *models.User, error)
    UpdateUser(user *models.User, password string) error
    DeleteUser(id uint) error
}

type HistoryService interface {
    FindByTaskID(taskId uint) ([]models.History, error)
    Restore(id uint) error
}
```

**职责：**
- 核心业务规则
- 业务验证（如：不能删除超管）
- 调用 DAO 进行数据操作
- 调用 PKG 服务（Git、Event、JWT）
- 事务管理（如需要）
- 事件发布

**关键特性：**
- ✅ 接口驱动设计（便于 Mock）
- ✅ 依赖注入（通过构造函数）
- ✅ 错误处理（返回 AppError）

---

### Controllers 模块

**文件结构：**
```
controllers/
├── user_controller.go       # 用户 HTTP 处理
├── repository_controller.go # 仓库 HTTP 处理
├── task_controller.go       # 任务 HTTP 处理
└── history_controller.go    # 历史版本 HTTP 处理
```

**标准方法结构：**
```go
func (ctrl *UserController) Login(c *gin.Context) {
    ctx := context.New(c)
    
    // 1. 参数绑定
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        ctx.SendErr(err)
        return
    }
    
    // 2. 调用 Service
    token, user, err := ctrl.userService.Login(req.Username, req.Password)
    if err != nil {
        ctrl.handleError(ctx, err)
        return
    }
    
    // 3. 返回响应
    ctx.SendOk("login successful", gin.H{"user": user, "token": token})
}
```

**职责：**
- HTTP 请求处理
- 参数绑定和验证
- 调用 Service 执行业务
- 统一响应格式
- 错误处理

**不包含：**
- ❌ 业务逻辑
- ❌ 数据库操作
- ❌ 复杂的数据转换

---

### PKG 模块

**文件结构：**
```
pkg/
├── errors/      # 统一错误定义
├── database/    # 数据库连接管理
├── event/       # SSE 事件系统
├── command/     # Shell 命令执行
├── context/     # Gin 上下文工具
├── jwt/         # JWT 认证
└── logger/      # 日志记录
```

**特点：**
- 独立的技术组件
- 可被任何层调用
- 无业务逻辑
- 可复用到其他项目

---

## 🔐 权限控制

### 角色定义

```go
const (
    UserRoleReporter   = 10  // 观察者
    UserRoleDeveloper  = 6   // 开发者
    UserRoleAdmin      = 3   // 管理员
    UserRoleSuperAdmin = 1   // 超级管理员
)
```

**角色说明：**
- **SuperAdmin (1)**: 超级管理员，拥有所有权限
- **Admin (3)**: 管理员，可管理用户和仓库
- **Developer (6)**: 开发者，可执行构建任务
- **Reporter (10)**: 观察者，仅可查看信息

**权限等级：** 数值越小，权限越高

### 权限中间件

```go
// 观察者及以上
v1.GET("/users", middlewares.RoleReporter(), userCtrl.Find)
v1.GET("/histories", middlewares.RoleReporter(), historyCtrl.FindByTaskID)

// 开发者及以上
v1.PUT("/repository", middlewares.RoleDeveloper(), repoCtrl.Edit)
v1.POST("/history/restore/:id", middlewares.RoleDeveloper(), historyCtrl.Restore)

// 仅超级管理员
v1.DELETE("/user/:id", middlewares.RoleSuperAdmin(), userCtrl.Destroy)
```

### 权限判断逻辑

```go
func role(c *gin.Context, roleVal int, errTip error) {
    ctx := context.New(c)
    claims, err := ctx.GetClaims()
    if err != nil {
        ctx.SendErr(err)
        c.Abort()
        return
    }
    
    // 用户角色值 <= 所需角色值 才有权限（数值越小，权限越高）
    if claims.Role > roleVal {
        ctx.SendErr(errTip)
        c.Abort()
        return
    }
    c.Next()
}
```

**权限等级说明：**
- SuperAdmin (1) 可访问所有接口
- Admin (3) 可访问 Admin 及以下权限的接口
- Developer (6) 可访问 Developer 及以下权限的接口
- Reporter (10) 仅可访问 Reporter 权限的接口

---

## 📡 事件系统

### 事件类型

**Repository 事件：**
- `cloneOk` - 克隆成功
- `cloneErr` - 克隆失败
- `editOk` - 编辑成功
- `destroyOk` - 删除成功

**Task 事件：**
- `runOk` - 开始运行
- `buildOk` - 构建成功
- `buildFail` - 构建失败
- `destroyOk` - 删除成功

### SSE 推送

```go
// 订阅事件
event.RepoSource.Subscribe(userID)

// 发布事件（排除自己）
event.RepoSource.PublishMsgExcludeID(eventType, data, userID)

// 发布事件（广播）
event.TaskSource.PublishMsg(eventType, data)
```

---

## 🧪 测试策略

### 1. 单元测试

**DAO 层：**
```go
func TestUserDAO_Create(t *testing.T) {
    // 使用内存数据库
    db := setupTestDB()
    dao := NewUserDAO(db)
    
    user := &models.User{Username: "test", Password: "123456"}
    err := dao.Create(user)
    
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}
```

**Service 层：**
```go
func TestUserService_Login(t *testing.T) {
    // Mock UserDAO
    mockDAO := &MockUserDAO{}
    mockDAO.On("FindByUsername", "test").Return(&models.User{...}, nil)
    
    service := NewUserService(mockDAO)
    token, user, err := service.Login("test", "password")
    
    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}
```

**Controller 层：**
```go
func TestUserController_Login(t *testing.T) {
    // Mock UserService
    mockService := &MockUserService{}
    mockService.On("Login", "test", "password").Return("token", &models.User{...}, nil)
    
    ctrl := NewUserController(mockService)
    
    // 创建测试请求
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/login", body)
    
    ctrl.Login(c)
    
    assert.Equal(t, 200, w.Code)
}
```

## 🎉 总结

MareWood 的新架构采用了业界成熟的三层架构模式，结合依赖注入和接口驱动设计，确保了：

✅ **高内聚**：每层职责单一明确
✅ **低耦合**：层与层之间通过接口通信
✅ **易测试**：可以轻松 Mock 任何层
✅ **易扩展**：添加新功能只需按层添加
✅ **易维护**：修改不会影响其他层

这为项目的长期发展奠定了坚实的基础。
