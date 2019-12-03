package serviceHelper

import (
	"FEDeployService/config"
	"os"
	"os/exec"
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

	if _, err := exec.LookPath("git"); err != nil {
		return err
	}

	if _, err := exec.LookPath("node"); err != nil {
		return err
	}

	for _,dependentTools := range config.Cfg.DependentTools{
		if _, err := exec.LookPath(dependentTools); err != nil {
			return err
		}
	}

	return nil
}


func initDir() error {

	if info, err := os.Stat(config.Cfg.ResourcesDir); err != nil || !info.IsDir(){
		if os.Mkdir(config.Cfg.ResourcesDir, os.ModePerm) != nil{
			return err
		}
	}

	if info, err := os.Stat(config.Cfg.RepositoryDir); err != nil || !info.IsDir(){
		if os.Mkdir(config.Cfg.RepositoryDir, os.ModePerm) != nil{
			return err
		}
	}
	if info, err := os.Stat(config.Cfg.WebRootDir); err != nil || !info.IsDir(){
		if os.Mkdir(config.Cfg.WebRootDir, os.ModePerm) != nil{
			return err
		}
	}
	return nil
}