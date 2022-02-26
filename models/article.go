package models

import "time"

type Article struct {
	Id         uint
	Title      string
	Keyword    string
	Author     string
	Desc       string
	Content    string
	Numbers    string
	Img_url    string
	Type       int
	State      int
	Tags       string
	CategoryId uint
	Category   string
	Views      int
	Comments   int
	Likes      int
	CreateTime time.Time
	UpdateTime time.Time
}
