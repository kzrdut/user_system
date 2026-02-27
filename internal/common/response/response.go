package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code int    `json:"code"` // 业务状态码
	Msg  string `json:"msg"`  // 提示信息
	Data any    `json:"data"` // 数据载体
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	})
}

// SuccessWithMsg 成功响应 (自定义消息)
func SuccessWithMsg(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeUnknownError,
		Msg:  msg,
		Data: nil,
	})
}

// FailWithCode 失败响应 (自定义业务码)
func FailWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
