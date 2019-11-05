package service

import (
	"log"
	"os/exec"
)

func CheckingEnv()  {

	if _, err := exec.LookPath("git"); err != nil {
		log.Fatalf("没有检测到 git 命令，请先安装 git !")
	}

	if _, err := exec.LookPath("node"); err != nil {
		log.Fatalf("没有检测到 node 命令，请先安装 node !")
	}

	if _, err := exec.LookPath("npm"); err != nil {
		log.Fatalf("没有检测到 npm 命令，请先安装 npm !")
	}

	if _, err := exec.LookPath("yarn"); err != nil {
		log.Fatalf("没有检测到 yarn 命令，请先安装 yarn !")
	}


}