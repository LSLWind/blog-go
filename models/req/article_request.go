package req

import "time"

/**
文章查询
*/
type ArticleQuery struct {
	Id         uint
	Title      string
	Keyword    string
	Author     string
	Desc       string
	Type       int
	State      int
	Tags       string
	CategoryId uint
	Category   string
	CreateTime time.Time
	UpdateTime time.Time
}

type ArticleRequest struct {
	Id         uint
	Title      string
	Keyword    string
	Author     string
	Desc       string
	Type       int
	State      int
	Tags       string
	CategoryId uint
	Category   string
	CreateTime time.Time
	UpdateTime time.Time
}

// 增加文章请求
type AddArticleRequest struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	CategoryId uint   `json:"categoryId"`
}
