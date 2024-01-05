package conf

import (
	"os"
	"path"
)

var Version, AppName, AppRepository string
var (
	DbDns         string
	ResourcesDir  string
	RepositoryDir string
	WebRootDir    string
	CurrentDir    string
	ClientDir     string
)
var DependTools = make(map[string]string)

//ResourcesDir 容器挂载目录，包含数据库文件、仓库和打包的静态文件目录

func init() {
	Version = "1.0.0"
	AppName = "MareWood"
	AppRepository = "https://github.com/xusenlin/MareWood"

	var err error
	if CurrentDir, err = os.Getwd(); err != nil {
		panic(err)
	}
	ClientDir = path.Join(CurrentDir, "ui", "dist")
	ResourcesDir = path.Join(CurrentDir, "resources")

	DbDns = path.Join(ResourcesDir, "database.db")
	RepositoryDir = path.Join(ResourcesDir, "repositories")
	WebRootDir = path.Join(ResourcesDir, "webs")

	DependTools = map[string]string{"git": "", "node": "", "npm": "", "yarn": "", "pnpm": ""}
}

func ConfigMap() map[string]any {
	return map[string]any{
		"version":       Version,
		"appName":       AppName,
		"appRepository": AppRepository,

		"dbDns":         DbDns,
		"resourcesDir":  ResourcesDir,
		"repositoryDir": RepositoryDir,
		"webRootDir":    WebRootDir,
		"clientDir":     ClientDir,

		"dependTools": DependTools,

		"httpPort": HttpPort,
		"ginMode":  GinMode,
	}
}
