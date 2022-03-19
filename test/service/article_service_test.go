package service

import (
	"blog-go/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryByPageAndSize(t *testing.T) {
	a := service.ArticleService{}
	assert.NotNil(t, a.QueryByPageAndSize(1, 10))
}
