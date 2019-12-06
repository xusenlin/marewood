package helper

import (
	"math/rand"
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
