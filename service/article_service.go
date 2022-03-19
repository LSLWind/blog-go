package service

import (
	"blog-go/dao"
	"blog-go/models"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type ArticleService struct {
	articleDao dao.ArticleDao
}

func (s ArticleService) FindAll() ([]models.Article, int64) {
	return s.articleDao.QueryAll()
}

//分页查询
func (a ArticleService) QueryByPageAndSize(page int, pageSize int) []models.Article {
	if page <= 0 || pageSize <= 0 {
		logger.Error("分页查询参数错误", page, pageSize)
	}
	return a.articleDao.QueryByPageAndSize(page, pageSize)
}
