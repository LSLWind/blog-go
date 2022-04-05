package utils

import (
	"blog-go/utils"
	"testing"
)

func TestArticleDirSync(t *testing.T) {
	a := utils.ArticleUtil{}
	a.ArticleDirSync("F:\\笔记文档\\研究生")
}

func TestMain(m *testing.M) {
	m.Run()
}
