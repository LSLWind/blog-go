package dao

import (
	"blog-go/models"
	"fmt"
)

type ArticleDao struct {
}

//查询请求
func (a ArticleDao) QueryAll() ([]models.Article, int64) {
	var resp []models.Article

	globalDb.Find(&resp)
	fmt.Println(resp)
	return resp, int64(len(resp))
}

func (a ArticleDao) Testq() {
	var article models.Article
	globalDb.First(&article)
	fmt.Println(article)
}
