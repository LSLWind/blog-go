package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

func initTagRouter(router *gin.RouterGroup) {
	tagApi := new(api.TagApi)
	tagRouter := router.Group("/")
	{
		tagRouter.GET("/getTagList", tagApi.GetTagList)
	}
}
