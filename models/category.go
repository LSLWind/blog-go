package models

import "time"

type Category struct {
	Id              uint
	Name            string
	SubCategoryId   int
	SubCategoryName string
	CreateTime      time.Time
	UpdateTime      time.Time
}
