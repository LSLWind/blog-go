package req

/**
  请求目录，以及关联的文章信息
*/
type CategoryRequest struct {
	Name                string
	SubCategoryNameList []string
}
type CategoryArticle struct {
	ArticleId   uint
	ArticleName string
}

/**
一级目录
*/
type CategoryArticleRequest struct {
	Name            string
	SubCategoryList []SecondCategory
	SubArticleList  []CategoryArticle
}

/**
二级目录
*/
type SecondCategory struct {
	Name           string
	SubArticleList []CategoryArticle
}

// 一级目录信息
type OneLevelCategoryRequest struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	SvgIcon  string `json:"svgIcon"`
}
