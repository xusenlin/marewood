package helper

import (
	"errors"
	"net/url"
	"strings"
)

//将git url 转换为带权限的url
func GitUrl2AuthUrl(gitUrl string, userName string, password string) (string, error) {

	urlArr := strings.SplitAfter(gitUrl, "://")
	if 2 != len(urlArr) {
		return "", errors.New("the address is not compliant!")
	}

	authUrl := urlArr[0] + url.PathEscape(userName) + ":" + url.PathEscape(password) + "@" + urlArr[1]

	return authUrl, nil
}

//通过git url 返回仓库的名字
func GetRepositoryNameByUrl(gitUrl string) (string, error) {

	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("the address must have a .git suffix! ")
	}
	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")

	urlArr := strings.Split(noSuffixUrl, "/")

	return urlArr[ len(urlArr) - 1 ], nil

}
