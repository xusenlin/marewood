package task

type taskStatus uint

const ( //任务执行状态
	StatusDefault taskStatus = iota
	StatusProcessing
	StatusFailed
	StatusSuccess
)
