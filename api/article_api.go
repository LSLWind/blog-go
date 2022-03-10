package api

import (
	"blog-go/models"
	"blog-go/service"
	"fmt"
	"github.com/gin-gonic/gin"
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
