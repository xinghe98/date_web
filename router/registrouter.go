package router

import (
	"datevisual/V2/controller"
	"github.com/gin-gonic/gin"
)

func RegistRoutes(r *gin.Engine) *gin.Engine {
	date := r.Group("/api/")
	{
		date.GET("ping", controller.Pong)
	}
	return r
}
