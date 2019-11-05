package service

import (
	"FEDeployService/config"
	"log"
	"os"
)

func InitDir() {

	if info, err := os.Stat(config.Cfg.RepositoryDir); err != nil || !info.IsDir(){
		if os.Mkdir(config.Cfg.RepositoryDir, os.ModePerm) != nil{
			log.Fatalf("创建“"+config.Cfg.RepositoryDir+"”目录失败！")
		}
	}
	if info, err := os.Stat(config.Cfg.WebRootDir); err != nil || !info.IsDir(){
		if os.Mkdir(config.Cfg.WebRootDir, os.ModePerm) != nil{
			log.Fatalf("创建“"+config.Cfg.RepositoryDir+"”目录失败！")
		}
	}
}
