package api

import (
	"blog-go/models/req"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
	articleService service.ArticleService
}

/**
查询文章列表
*/
func (a ArticleApi) Find(c *gin.Context) {
	query := req.ArticleQuery{}

	if c.Bind(&query) == nil {
		list, _ := a.articleService.FindAll(query)
		c.JSON(200, list)
	}
}
