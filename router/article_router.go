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
	articleRouter := router.Group("/articles")

	{
		articleRouter.GET("/getArticleList", articleApi.GetArticleList)
		articleRouter.GET("/getArticleListByPageAndSize", articleApi.GetArticleListByPageAndSize)
		articleRouter.GET("/getArticleById", articleApi.GetArticleById)
		articleRouter.POST("/addOneArticle", articleApi.AddOneArticle)
	}

}
