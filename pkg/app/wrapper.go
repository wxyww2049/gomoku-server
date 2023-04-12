package app

import (
	"github.com/gin-gonic/gin"
)

/**
对于返回结果进行一层包装
*/

type Result struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	wrapper *Wrapper
}

func (r Result) SendJSON() {
	r.wrapper.Ctx.JSON(200, r)
}

type Wrapper struct {
	Ctx *gin.Context
}

func NewWrapper(c *gin.Context) *Wrapper {
	return &Wrapper{Ctx: c}
}

func (w Wrapper) OK() Result {
	return Result{
		Code:    0,
		Msg:     "",
		Data:    nil,
		wrapper: &w,
	}
}

func (w Wrapper) Success(data interface{}) Result {
	return Result{
		Code:    0,
		Msg:     "",
		Data:    data,
		wrapper: &w,
	}
}
func (w Wrapper) Error(msg string) Result {
	return Result{
		Code:    -1,
		Msg:     msg,
		Data:    nil,
		wrapper: &w,
	}
}
func (w Wrapper) ErrorWithCode(code int, msg string) Result {
	return Result{
		Code:    code,
		Msg:     msg,
		Data:    nil,
		wrapper: &w,
	}
}
func (w Wrapper) GetIP() string {
	return w.Ctx.ClientIP()
}
