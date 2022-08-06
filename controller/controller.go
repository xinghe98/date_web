package controller

import (
	"datevisual/V2/service"
	"github.com/gin-gonic/gin"
)

// Pong 测试路由的控制器
func Pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "pong"})
}

// CreateData 用于新建当日数据的控制器
func CreateData(ctx *gin.Context) {
	service.Create(ctx)
}
