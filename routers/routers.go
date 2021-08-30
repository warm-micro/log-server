package routers

import (
	"wm/logger/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("log/ping", controllers.Pong)
	AddLoggerRouter(router.Group("log"))
}
