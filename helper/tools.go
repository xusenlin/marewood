package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//随机字符串
func RandSeq(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func DigestString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func BuildDependCmd(tools string) (name string, arg string) {
	switch tools {
	case "npm":
		return "npm", "install"
	case "yarn":
		return "yarn", ""
	case "cnpm":
		return "cnpm", "install"
	default:
		return "npm", "install"
	}
}
func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}


func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func RemoveDir(dir string) error {

	if !IsDir(dir) {
		return  errors.New("找不到目录,无法执行删除")
	}

	if err := os.RemoveAll(dir);err != nil {
		return err
	}
	return  nil
}