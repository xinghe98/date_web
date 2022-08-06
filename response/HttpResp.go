// Package HttpResp 封装响应方法，避免过多大括号
package HttpResp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Res 其他响应
func Res(ctx *gin.Context, status int, code int, data interface{}, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// ResOK 请求成功的响应。
func ResOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}
