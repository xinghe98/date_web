package router

import (
	"datevisual/V2/controller"
	"github.com/gin-gonic/gin"
)

// RegistRoutes 注册路由
func RegistRoutes(r *gin.Engine) *gin.Engine {
	date := r.Group("/api/")
	{
		date.GET("ping", controller.Pong)
		date.POST("data", controller.CreateData)
		date.GET("data", controller.FindDate)
	}
	newuser := r.Group("/api/")
	{
		newuser.POST("newusers", controller.CreateUsers)
		newuser.GET("newusers", controller.FindAllUsers)
	}
	return r
}
