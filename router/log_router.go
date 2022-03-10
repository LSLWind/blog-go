package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

func initLogRouter(router *gin.RouterGroup) {
	logApi := new(api.LogApi)
	logRouter := router.Group("/logs")
	{
		logRouter.GET("/getAllLog", logApi.GetAllLogsContent)
	}
}
