package response

import "time"

//gen --connstr "root:123456@tcp(127.0.0.1:3306)/blog_go?&parseTime=True" --database articles  --json --gorm --from --guregu
type ArticleResponse struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Keyword    string    `json:"keyword"`
	Author     string    `json:"author"`
	Desc       string    `json:"desc"`
	Content    string    `json:"content"`
	Numbers    string    `json:"numbers"`
	ImgUrl     string    `json:"imgUrl"`
	Type       int       `json:"type"`
	State      int       `json:"state"`
	Tags       string    `json:"tags"`
	CategoryId uint      `json:"categoryId"`
	Category   string    `json:"category"`
	Views      int       `json:"views"`
	Comments   int       `json:"comments"`
	Likes      int       `json:"likes"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime string    `json:"updateTime"`
}
