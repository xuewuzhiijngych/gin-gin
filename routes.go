package main

import (
	"github.com/gin-gonic/gin"
	"ych.hub/controller"
	"ych.hub/middleware"
)

// CollectRoute 注册路由
func CollectRoute(app *gin.Engine) *gin.Engine {
	app.POST("/api/user/register", controller.Register)
	app.POST("/api/user/login", controller.Login)
	app.GET("/api/user/info", middleware.AuthMiddleware(), controller.Info)
	return app
}
