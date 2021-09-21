package routers

import (
	"wm/logger/controllers"

	"github.com/gin-gonic/gin"
)

func AddLoggerRouter(router *gin.RouterGroup) {
	router.GET("", controllers.ListLogs)
	router.POST("", controllers.Addlog)

	router.GET("/counts", controllers.ListCounts)

	router.GET("/deleteWrong", controllers.DeleteWrongLog)
}
