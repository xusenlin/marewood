package initial

import (
	"marewood/conf"
	"os"
	"os/exec"
	"strings"
)

func CheckEnvAndInitDir() error {
	if err := checkEnv(); err != nil {
		return err
	}
	if err := initDir(); err != nil {
		return err
	}
	return nil
}

func checkEnv() error {

	for name, _ := range conf.DependTools {
		if _, err := exec.LookPath(name); err != nil {
			return err
		}

		cmd := exec.Command(name, "--version")
		output, err := cmd.Output()
		if err == nil {
			conf.DependTools[name] = strings.TrimSpace(string(output))
		}
	}
	return nil
}

func initDir() error {

	if info, err := os.Stat(conf.ResourcesDir); err != nil || !info.IsDir() {
		if os.Mkdir(conf.ResourcesDir, os.ModePerm) != nil {
			return err
		}
	}

	if info, err := os.Stat(conf.RepositoryDir); err != nil || !info.IsDir() {
		if os.Mkdir(conf.RepositoryDir, os.ModePerm) != nil {
			return err
		}
	}
	if info, err := os.Stat(conf.WebRootDir); err != nil || !info.IsDir() {
		if os.Mkdir(conf.WebRootDir, os.ModePerm) != nil {
			return err
		}
	}
	return nil
}
