package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseResult struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	TimeStamp int64       `json:"timeStamp"`
}

func Success(data interface{}) *ResponseResult {
	msg := &ResponseResult{
		Code:      SuccessCode,
		Msg:       Text(SuccessCode),
		Data:      data,
		TimeStamp: time.Now().UnixMilli(),
	}
	return msg
}

func Fail(msg string) *ResponseResult {
	msgObj := &ResponseResult{
		Code:      http.StatusInternalServerError,
		Msg:       msg,
		TimeStamp: time.Now().UnixMilli(),
	}
	return msgObj
}

func FailMessage(code int, msg string) *ResponseResult {
	msgObj := &ResponseResult{
		Code:      code,
		Msg:       msg,
		TimeStamp: time.Now().UnixMilli(),
	}
	return msgObj
}

func FailMessageGin(code int, msg string, c *gin.Context) *ResponseResult {
	msgObj := &ResponseResult{
		Code:      code,
		Msg:       msg,
		TimeStamp: time.Now().UnixMilli(),
	}
	c.JSON(http.StatusOK, msgObj)
	return msgObj
}
