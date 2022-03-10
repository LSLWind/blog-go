package service

import (
	"blog-go/dao"
	"blog-go/models"
)

type ArticleService struct {
	articleDao dao.ArticleDao
}

func (s ArticleService) FindAll() ([]models.Article, int64) {
	return s.articleDao.QueryAll()
}
