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
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func RemoveDir(dir string) error {

	if !IsDir(dir) {
		return errors.New("no directory found, cannot be deleted")
	}

	return os.RemoveAll(dir)
}

func InStrArr(str string, arr []string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func ComputeTotalPage(total, pageSize int, ) int {
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	return totalPage
}
