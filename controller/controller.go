package controller

import "github.com/gin-gonic/gin"

func Pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "pong"})
}
