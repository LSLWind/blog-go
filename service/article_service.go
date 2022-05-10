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

//根据文章id查询
func (a ArticleService) QueryArticleById(id int) (articleContent response.ArticleContentResponse) {
	article := a.articleDao.QueryArticleById(id)
	logger.Info("查询文章，id = {", id, "}，title = {", article.Title, "}")
	return article2ArticleContentResponse(article)
}

/**
文章内容返回数据响应
- 返回全部信息
- 时间格式化
*/
func article2ArticleContentResponse(article models.Article) (articleContent response.ArticleContentResponse) {
	articleContent.Id = article.Id
	articleContent.Title = article.Title
	articleContent.Keyword = article.Keyword
	articleContent.Author = article.Author
	articleContent.Desc = article.Desc
	articleContent.Content = article.Content
	articleContent.Numbers = article.Numbers
	articleContent.ImgUrl = article.ImgUrl
	articleContent.Type = article.Type
	articleContent.Tags = article.Tags
	articleContent.CategoryId = article.CategoryId
	articleContent.Category = article.Category
	articleContent.Views = article.Views
	articleContent.Comments = article.Comments
	articleContent.Likes = article.Likes
	timeUtil := utils.TimeUtil{}
	articleContent.CreateTime = timeUtil.GetYearMonthDay(article.CreateTime)
	articleContent.UpdateTime = timeUtil.GetYearMonthDay(article.UpdateTime)
	return articleContent
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
		articleResponse.ImgUrl = a.ImgUrl
		articleResponse.Views = a.Views
		articleResponse.Comments = a.Comments
		articleResponse.Likes = a.Likes
		timeUtil := utils.TimeUtil{}
		articleResponse.UpdateTime = timeUtil.GetYearMonthDay(a.UpdateTime)
		articleResponses[i] = articleResponse
	}
	return articleResponses
}
