package dto

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewMsg(code int, data interface{}) *Message {
	return &Message{Code: code, Data: data}
}

func NewErrMsg(err error) *Message {
	return NewMsg(-1, err.Error())
}
