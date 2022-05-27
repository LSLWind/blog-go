package models

import "time"

// svg图标信息
type SvgIcon struct {
	Id         uint
	Name       string
	SvgStr     string
	Size       uint
	CreateTime time.Time
	UpdateTime time.Time
}
