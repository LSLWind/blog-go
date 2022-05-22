package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

/**
初始化gin 路由映射
*/

func Init() *gin.Engine {
	router := gin.New()
	//可用 TODO
	router.Use(gin.Logger())
	//	router.Use(gin.Recovery())

	// 基于cookie创建session的存储引擎，传递一个参数，用来做加密时的密钥
	store := cookie.NewStore([]byte("baopowodedoushisb"))
	//session中间件生效，参数mySession，是浏览器端cookie的名字
	router.Use(sessions.Sessions("currentUser", store))

	//api
	apiRouter := router.Group("/api")
	{
		initArticleRouter(apiRouter)
		initTagRouter(apiRouter)
		initLogRouter(apiRouter)
		initNavigationRouter(apiRouter)
		initCategoryRouter(apiRouter)
		initLslRouter(apiRouter)
	}

	return router
}
