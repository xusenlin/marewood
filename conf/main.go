package conf

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"syscall"
)

var (
	DbDns         string
	ResourcesDir  string
	RepositoryDir string
	WebRootDir    string
	CurrentDir    string
	ClientDir     string
)
var DependTools = make(map[string]string)

var Version, AppName, AppRepository, SystemInfo string

func byteCountToHuman(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(b)/float64(div), "KMGTPE"[exp])
}

//ResourcesDir 容器挂载目录，包含数据库文件、仓库和打包的静态文件目录

func init() {
	Version = "1.2.0"
	AppName = "MareWood"
	AppRepository = "https://github.com/xusenlin/MareWood"

		SystemInfo = fmt.Sprintf("%s / %s", runtime.GOOS, runtime.GOARCH)

		var localErr error
	CurrentDir, localErr = os.Getwd()
	if localErr != nil {
		panic(localErr)
	}
	ClientDir = path.Join(CurrentDir, "ui", "dist")
	ResourcesDir = path.Join(CurrentDir, "resources")

	DbDns = path.Join(ResourcesDir, "database.db")
	RepositoryDir = path.Join(ResourcesDir, "repositories")
	WebRootDir = path.Join(ResourcesDir, "webs")

	DependTools = map[string]string{"git": "", "node": "", "npm": "", "yarn": "", "pnpm": ""}
}

func ConfigMap() map[string]any {
	osInfo := SystemInfo // Get the base OS info from the init function

	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	var diskUsage string
	if err != nil {
		diskUsage = "N/A"
	} else {
		all := stat.Blocks * uint64(stat.Bsize)
		free := stat.Bfree * uint64(stat.Bsize)
		used := all - free
		diskUsage = fmt.Sprintf("Total:%s, Used: %s, Free: %s", byteCountToHuman(all), byteCountToHuman(used), byteCountToHuman(free))
	}

	// Combine OS info and dynamic disk usage
	combinedSystemInfo := fmt.Sprintf("%s, Disk: %s", osInfo, diskUsage)

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
		"systemInfo": combinedSystemInfo,
	}
}
