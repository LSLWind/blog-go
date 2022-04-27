package api

import (
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type CategoryApi struct {
	categoryService service.CategoryService
}

func (categoryApi CategoryApi) GetAllCategoryInfo(c *gin.Context) {
	categoryResponses := categoryApi.categoryService.GetAllCategory()
	resp.OK(c, categoryResponses)
}
