package models

import "time"

type Article struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Keyword    string    `json:"keyword"`
	Author     string    `json:"author"`
	Desc       string    `json:"desc"`
	Content    string    `json:"content"`
	Numbers    string    `json:"numbers"`
	ImgUrl     string    `json:"imgUrl"`
	Tags       string    `json:"tags"`
	CategoryId uint      `json:"categoryId"`
	Category   string    `json:"category"`
	Views      int       `json:"views"`
	Comments   int       `json:"comments"`
	Likes      int       `json:"likes"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
