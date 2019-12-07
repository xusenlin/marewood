package config

type WebHooksConfig struct {
	EventKey     string
	PushEventVal string
	SignatureKey string
}

var GitlabPushEvent = WebHooksConfig{"X-Gitlab-Event", "Push Hook", "X-Gitlab-Token"}
var GithubPushEvent = WebHooksConfig{"X-GitHub-Event", "push", "X-Hub-Signature"}
