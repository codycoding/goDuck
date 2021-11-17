package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//
// Result
//  @Description: 基本返回结构
//  @receiver h
//  @param code
//  @param data
//  @param msg
//  @param c
//
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

//
// SuccessWithMessage
//  @Description: 成功且回复信息
//  @receiver h
//  @param message
//  @param c
//
func SuccessWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, c)
}

//
// SuccessWithData
//  @Description: 成功且返回数据
//  @receiver h
//  @param data
//  @param c
//
func SuccessWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "操作成功", c)
}

//
// FailWithMessage
//  @Description: 失败且返回信息
//  @receiver h
//  @param code
//  @param message
//  @param c
//
func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

//
// FailWithData
//  @Description: 失败且返回数据
//  @receiver h
//  @param code
//  @param data
//  @param c
//
func FailWithData(code int, data interface{}, c *gin.Context) {
	Result(code, data, "操作失败", c)
}

//
// UnauthorizedWithMessage
//  @Description: 授权错误返回信息
//  @param message
//  @param c
//
func UnauthorizedWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Msg: message,
	})
}
