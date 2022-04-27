package models

import "time"

type Category struct {
	Id              uint      `json:"id"`
	Name            string    `json:"name"`
	SubCategoryId   uint      `json:"subCategoryId"`
	SubCategoryName string    `json:"subCategoryName"`
	CreateTime      time.Time `json:"createTime"`
	UpdateTime      time.Time `json:"updateTime"`
}
