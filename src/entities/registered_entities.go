package entities

import (
	"news-go/src/entities/article_categories"
	"news-go/src/entities/articles"
	"news-go/src/entities/categories"
)

var RegisteredEntities = []any{
	categories.Category{},
	articles.Article{},
	article_categories.ArticleCategory{},
}
