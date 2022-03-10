package router

import (
	"blog-go/api"
	"github.com/gin-gonic/gin"
)

//gin router 组
/**
路由组，映射url与api方法
*/
func initArticleRouter(router *gin.RouterGroup) {
	articleApi := new(api.ArticleApi)
	articleRouter := router.Group("/")

	{
		articleRouter.GET("/getArticleList", articleApi.GetArticleList)
	}

}
