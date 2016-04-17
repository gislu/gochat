package utils

// Entry 结构定了最基本的条目结构
type Entry struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}
