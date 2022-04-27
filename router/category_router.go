package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

//gin router 组
/**
路由组，映射url与api方法
*/
func initCategoryRouter(router *gin.RouterGroup) {
	categoryApi := new(api.CategoryApi)
	categoryRouter := router.Group("/category")

	{
		categoryRouter.GET("/getCategoryList", categoryApi.GetAllCategoryInfo)
	}

}
