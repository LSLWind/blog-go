package api

import (
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type SvgIconApi struct {
	svgIconService service.SvgIconService
}

// 查询svg icon信息 url/?name =
func (s SvgIconApi) GetSvgIconInfoByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		resp.Error(c, "参数错误：name={ "+name+" }")
	}
	// TODO 所有查询校验查询结果
	svgIconInfo := s.svgIconService.GetSvgIconByName(name)
	resp.OK(c, svgIconInfo)
}
