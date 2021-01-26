package main

import (
	"github.com/gin-gonic/gin"
	"ych.hub/controller"
)

// CollectRoute 注册路由
func CollectRoute(app *gin.Engine) *gin.Engine {
	app.POST("/api/user/register", controller.Register)
	return app
}
