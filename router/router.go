package router

import "github.com/gin-gonic/gin"

/**
初始化gin 路由映射
*/
func Init() *gin.Engine {
	router := gin.New()
	//可用 TODO
	router.Use(gin.Logger())
	//	router.Use(gin.Recovery())

	//api
	apiRouter := router.Group("/api")
	{
		initArticleRouter(apiRouter)
		initTagRouter(apiRouter)
		initLogRouter(apiRouter)
		initNavigationRouter(apiRouter)
		initCategoryRouter(apiRouter)
	}

	return router
}
