package dao

import "blog-go/models"

type SvgIconDao struct {
}

// 查询svg icon
func (s SvgIconDao) QuerySvgIconByName(name string) (svgIcon models.SvgIcon) {
	//where 条件查询
	globalDb.Table("svg_icon").Find(&svgIcon, "name = ?", name)
	return svgIcon
}
