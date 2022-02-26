package response

import "time"

//gen --connstr "root:123456@tcp(127.0.0.1:3306)/blog_go?&parseTime=True" --database articles  --json --gorm --from --guregu
type ArticleResponse struct {
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
