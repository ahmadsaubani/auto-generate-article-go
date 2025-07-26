package article_repositories

import (
	"context"
	"fmt"
	"news-go/src/configs/database"
	"news-go/src/entities/articles"
	"news-go/src/entities/categories"
	"news-go/src/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

type ArticleRepository interface {
	GetAll(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, error)
	SaveArticles(ctx context.Context, articlesList []articles.Article) error
	FindCategoryByValue(value string) (*categories.Category, error)
}

type articleRepository struct{}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

func (r *articleRepository) GetAll(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, error) {
	var articleList []articles.Article

	db := database.GormDB.
		Preload("Categories")
	err := helpers.GetAllModelsWithDB(ctx, db, &articleList, pag)

	return articleList, err
}

func (r *articleRepository) SaveArticles(ctx context.Context, articlesList []articles.Article) error {
	for _, article := range articlesList {
		_, err := r.FindArticleByURL(article.URL)
		if err == nil {
			// Jika tidak error berarti artikel ditemukan → skip
			continue
		}

		var articleCategories []categories.Category
		for _, name := range article.CategoryNames {
			ValueName := strings.ToLower(name) // ubah ke lowercase
			CapitalName := strings.ToUpper(string(ValueName[0])) + ValueName[1:]

			category, err := r.FindCategoryByValue(ValueName)
			if err != nil {
				// Jika tidak ditemukan, buat kategori baru
				newCategory := categories.Category{
					Label: CapitalName,
					Value: ValueName,
				}
				if err := helpers.InsertModel(&newCategory); err != nil {
					return fmt.Errorf("could not create category: %w", err)
				}
				category = &newCategory
			}
			articleCategories = append(articleCategories, *category)
		}
		article.Categories = articleCategories

		if err := helpers.InsertModel(&article); err != nil {
			return fmt.Errorf("could not insert article: %w", err)
		}
	}
	return nil
}

func (r *articleRepository) FindCategoryByValue(value string) (*categories.Category, error) {
	var category categories.Category

	err := helpers.FindOneByField(&category, "value", value)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}
	return &category, nil
}

func (r *articleRepository) FindArticleByURL(value string) (*articles.Article, error) {
	var article articles.Article

	err := helpers.FindOneByField(&article, "url", value)
	if err != nil {
		return nil, fmt.Errorf("article not found: %w", err)
	}
	return &article, nil
}
