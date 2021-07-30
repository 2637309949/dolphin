package types

// Error defined TODO
type Error struct {
	// 编码
	Code int `json:"code" xml:"code"`
	// 消息
	Msg string `json:"msg" xml:"msg"`
}

// Error defined error helpers
func (e Error) Error() string {
	return e.Msg
}
