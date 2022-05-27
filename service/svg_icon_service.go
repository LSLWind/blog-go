package service

import (
	"blog-go/dao"
	"blog-go/models/response"
)

type SvgIconService struct {
	svgIconDao dao.SvgIconDao
}

// 获取svgIcon信息
func (s SvgIconService) GetSvgIconByName(name string) (svgIconInfo response.SvgIconInfo) {
	svgIcon := s.svgIconDao.QuerySvgIconByName(name)
	svgIconInfo.Name = svgIcon.Name
	svgIconInfo.SvgStr = svgIcon.SvgStr
	svgIconInfo.Size = svgIcon.Size
	return svgIconInfo
}
