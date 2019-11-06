package service

import (
	"errors"
	"net/url"
	"strings"
)

func GitUrl2AuthUrl(gitUrl string, userName string, password string) (string, error) {

	urlArr := strings.SplitAfter(gitUrl, "://")
	if 2 != len(urlArr) {
		return "", errors.New("git url 地址不合规！")
	}

	authUrl := urlArr[0] + url.PathEscape(userName) + ":" + url.PathEscape(password) + "@" + urlArr[1]

	return authUrl, nil
}

func GetRepositoryNameByUrl(gitUrl string) (string, error) {

	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("git url 地址必须以 .git 结尾！")
	}
	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")
	urlArr := strings.Split(noSuffixUrl, "/")

	return urlArr[ len(urlArr) - 1 ], nil

}
