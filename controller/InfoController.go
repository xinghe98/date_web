package controller

import (
	"datevisual/V2/service"
	"github.com/gin-gonic/gin"
)

func CreateUsers(ctx *gin.Context) {
	service.CreateUser(ctx)
}
