package models

import "time"

//导航书签信息
type Navigation struct {
	Id         uint
	Type       string
	LinkUrl    string
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}
