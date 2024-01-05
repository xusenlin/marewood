package event

import eventSource "github.com/xusenlin/event-source"

const TaskTypeEditOk = "editOk"
const TaskTypeDestroyOk = "destroyOk"

// repository
const TaskTypeRunOk = "runOk"
const TaskTypeBuildOk = "buildOk"
const TaskTypeBuildFail = "buildFail"

type TaskData struct {
	TaskId   uint   `json:"taskId"`
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Msg      string `json:"msg"`
}

var TaskSource = eventSource.New[*TaskData, uint]()
