package command

import (
	"errors"
	"github.com/xusenlin/command"
	"net/url"
	"strings"
)

func gitUrl2AuthUrl(gitUrl string, userName string, password string) (string, error) {

	urlArr := strings.SplitAfter(gitUrl, "://")
	if 2 != len(urlArr) {
		return "", errors.New("the address is not compliant!")
	}

	authUrl := urlArr[0] + url.PathEscape(userName) + ":" + url.PathEscape(password) + "@" + urlArr[1]

	return authUrl, nil
}

func GitClone(userName, password, defaultBranch, url, dir string) (string, error) {
	var gitUrl string
	var err error
	if userName == "" || password == "" {
		gitUrl = url
	} else {
		gitUrl, err = gitUrl2AuthUrl(url, userName, password)
		if err != nil {
			return "", err
		}
	}
	return command.New("git").AddArgs("clone", "-b", defaultBranch, gitUrl, dir).Run()
}

func GitPull(dir string) (string, error) {
	cmd := command.New("git").AddArgs("pull")
	return cmd.RunInDir(dir)
}

func GitBranchList(dir string) ([]string, error) {
	out, err := command.New("git").AddArgs("branch", "-r").RunInDir(dir)
	if err != nil {
		return []string{}, errors.New(out)
	}
	deleteOrigin := strings.ReplaceAll(out, "origin/", "")

	branch := strings.Split(strings.Trim(strings.ReplaceAll(deleteOrigin, " ", ""), "\n"), "\n")

	var newBranch []string
	for _, b := range branch {
		if !strings.Contains(b, "HEAD") {
			newBranch = append(newBranch, b)
		}
	}
	return newBranch, nil
}

func GitPrune(dir string) (string, error) {
	return command.New("git").AddArgs("remote", "prune", "origin").RunInDir(dir)
}

func GitCheckoutDot(dir string) (string, error) {
	return command.New("git").AddArgs("checkout", ".").RunInDir(dir)
}

func GitCheckout(dir string, branch string) (string, error) {
	return command.New("git").AddArgs("checkout", branch).RunInDir(dir)
}

func GitHash(dir string) (string, error) {
	hash, err := command.New("git").AddArgs("log", "-1", "--pretty=format:%h").RunInDir(dir)
	if err != nil {
		return "", err
	}
	h := strings.TrimSpace(hash)
	return h, nil
}
