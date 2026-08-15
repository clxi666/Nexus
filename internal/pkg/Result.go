package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Response(c *gin.Context, Code int, msg string, data any) {
	c.JSON(Code,
		Result{
			Msg:  msg,
			Data: data,
		})
}

func OK(c *gin.Context, msg string, data any) {
	Response(c, http.StatusOK, msg, data)
}
func OKWithMsg(c *gin.Context, msg string) {
	OK(c, msg, gin.H{})
}
func OKWithData(c *gin.Context, data any) {
	OK(c, "请求成功", data)
}

func Fail(c *gin.Context, msg string) {
	Response(c, http.StatusInternalServerError, msg, gin.H{})
}
func FailWithCode(c *gin.Context, code int, msg string) {
	Response(c, code, msg, gin.H{})
}
