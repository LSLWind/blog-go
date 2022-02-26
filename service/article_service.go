package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/req"
)

type ArticleService struct {
	articleDao dao.ArticleDao
}

func (s ArticleService) FindAll(articleQuery req.ArticleQuery) ([]models.Article, int64) {
	return s.articleDao.QueryAll(articleQuery)
}
