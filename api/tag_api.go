package api

import (
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type TagApi struct {
	tagService service.TagService
}

func (a TagApi) GetTagList(c *gin.Context) {
	list, _ := a.tagService.QueryAll()
	c.JSON(200, list)
}
