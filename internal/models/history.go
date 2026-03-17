package models

// History 历史版本模型
type History struct {
	Model
	TaskID   uint   `json:"taskId"`
	Name     string `json:"name"`
	GitHash  string `json:"gitHash"`
	UserID   uint   `json:"userId"`
	UserName string `json:"userName"`
}

// TableName 指定表名
func (History) TableName() string {
	return "histories"
}
