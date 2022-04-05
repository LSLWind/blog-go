package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

func initNavigationRouter(router *gin.RouterGroup) {
	navigationApi := new(api.NavigationApi)
	NavigationRouter := router.Group("/navigations")
	{
		NavigationRouter.GET("/getAllNavigationInfo", navigationApi.GetAllNavigationInfo)
	}
}
