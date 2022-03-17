package utils

import (
	"blog-go/utils"
	"testing"
)

func TestArticleDirSync(t *testing.T) {
	a := utils.ArticleUtil{}
	a.ArticleDirSync("F:\\笔记文档\\迁移副本\\研究生\\c++\\makefile")
}

func TestMain(m *testing.M) {
	m.Run()
}
