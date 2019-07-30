
package api
// Reply api的回复结构
type Reply struct {
	Code    int         "json:\"code\""
	Message string      "json:\"message\""
	Data    interface{} "json:\"data\""
}

// NewReplyOk 默认的正确回复
func NewReplyOk() *Reply {
	return &Reply{
		Code:    1,
		Message: "成功",
		Data:    make([]struct{}, 0),
	}
}

// NewReplyError 默认的错误回复
func NewReplyError(msg string) *Reply {
	return &Reply{
		Code:    0,
		Message: msg,
		Data:    make([]struct{}, 0),
	}
}

