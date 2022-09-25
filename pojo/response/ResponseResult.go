package response

import (
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Result {
	msg := &Result{
		Code: http.StatusOK,
		Msg:  "请求成功",
		Data: data,
	}
	return msg
}

func Fail(msg string) *Result {
	msgObj := &Result{
		Code: http.StatusInternalServerError,
		Msg:  msg,
	}
	return msgObj
}

func FailMessage(code int, msg string) *Result {
	msgObj := &Result{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}
