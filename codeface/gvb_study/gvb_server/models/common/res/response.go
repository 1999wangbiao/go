package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, msg string, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OKWithMsg(data any, msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(Success, map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "位置错误", c)
}
