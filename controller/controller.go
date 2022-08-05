package controller

import (
	"datevisual/V2/service"
	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "pong"})
}
func CreateData(ctx *gin.Context) {
	service.Create(ctx)
}
