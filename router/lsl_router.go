package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

func initLslRouter(router *gin.RouterGroup) {
	lslApi := new(api.LslApi)
	lslRouter := router.Group("/lsl")
	{
		lslRouter.POST("/check", lslApi.LslCheckout)
		lslRouter.POST("/addOneLevelCategory", lslApi.AddOneLevelCategory)
	}
}
