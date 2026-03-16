package event

import eventSource "github.com/xusenlin/event-source"

const RepoTypeCloneOk = "cloneOk"
const RepoTypeCloneErr = "cloneErr"
const RepoTypeEditOk = "editOk"
const RepoTypeDestroyOk = "destroyOk"

type RepoData struct {
	RepositoryId uint   `json:"repositoryId"`
	UserId       uint   `json:"userId"`
	UserName     string `json:"userName"`
	Msg          string `json:"msg"`
}

var RepoSource = eventSource.New[*RepoData, uint]()
