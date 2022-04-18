package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/response"
	"blog-go/utils"
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
func (a ArticleService) QueryByPageAndSize(page int, pageSize int) []response.ArticleResponse {
	if page <= 0 || pageSize <= 0 {
		logger.Error("分页查询参数错误", page, pageSize)
	}
	resp := a.articleDao.QueryByPageAndSize(page, pageSize)
	return article2ArticleResponse(resp)
}

/**
文章返回数据响应
- 时间格式化
*/
func article2ArticleResponse(articles []models.Article) []response.ArticleResponse {
	articleResponses := make([]response.ArticleResponse, len(articles))
	for i, a := range articles {
		var articleResponse response.ArticleResponse
		articleResponse.Id = a.Id
		articleResponse.Title = a.Title
		articleResponse.Desc = a.Desc
		articleResponse.Numbers = a.Numbers
		articleResponse.Category = a.Category
		articleResponse.Views = a.Views
		articleResponse.Comments = a.Comments
		articleResponse.Likes = a.Likes
		timeUtil := utils.TimeUtil{}
		articleResponse.UpdateTime = timeUtil.GetYearMonthDay(a.UpdateTime)
		articleResponses[i] = articleResponse
	}
	return articleResponses
}
