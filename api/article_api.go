package api

import (
	"blog-go/models"
	"blog-go/models/req"
	"blog-go/resp"
	"blog-go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleApi struct {
	articleService service.ArticleService
}

type data struct {
	List  []models.Article `json:"list"`
	Count int64            `json:"count"`
}

/**
查询文章列表
*/
func (a ArticleApi) GetArticleList(c *gin.Context) {

	listRes, _ := a.articleService.FindAll()

	datares := data{
		List:  listRes,
		Count: 1,
	}
	fmt.Print(datares)
	c.JSON(200, gin.H{"code": 200, "message": "a", "data": listRes})
	//c.JSON(200,gin.H{"code":200,"message":"a","data":data{
	//	List: listRes,
	//	Count: 1,
	//}})

}

//分页查询，url/?key=v的形式
func (a ArticleApi) GetArticleListByPageAndSize(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	res := a.articleService.QueryByPageAndSize(page, pageSize)
	resp.OK(c, res)

}

//根据id查询文章，url/?id=
func (a ArticleApi) GetArticleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	articleResponse := a.articleService.QueryArticleById(id)
	resp.OK(c, articleResponse)
}

//增加一篇文章，json
func (a ArticleApi) AddOneArticle(c *gin.Context) {
	//绑定json格式数据
	articleRequest := req.AddArticleRequest{}

	if err := c.BindJSON(&articleRequest); err != nil {
		resp.Error(c, "参数错误")
	}

	if a.articleService.AddOneArticle(articleRequest) {
		resp.OK(c, "添加文章成功")
	} else {
		resp.Error(c, "添加文章失败")
	}

}
