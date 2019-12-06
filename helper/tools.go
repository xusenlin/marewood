package helper

import (
	"math/rand"
	"os"
	"strings"
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

func GetDependTools() []string {
	toolsStr := os.Getenv("DEPENDENT_TOOLS")

	if !strings.Contains(toolsStr,"|") {
		return []string{"npm"}
	}
	tools := strings.Split(toolsStr,"|")

	return tools
}