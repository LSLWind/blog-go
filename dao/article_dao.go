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

	return resp, int64(len(resp))
}

//分页查询
func (a ArticleDao) QueryByPageAndSize(page int, pageSize int) []models.Article {
	var resp []models.Article
	if err := globalDb.Limit(pageSize).Offset((page - 1) * pageSize).Order("update_time desc").Find(&resp).Error; err != nil {
		logger.Error("分页查询错误，{},{},{}", page, pageSize, err)
	}
	return resp
}

//根据id返回文章信息
func (a ArticleDao) QueryArticleById(id int) (article models.Article) {
	//where 条件查询
	globalDb.Find(&article, "id = ?", id)
	return article
}

func (a ArticleDao) Testq() {
	var article models.Article
	globalDb.First(&article)
	fmt.Println(article)
}
