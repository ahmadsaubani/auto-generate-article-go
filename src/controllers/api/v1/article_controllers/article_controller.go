package article_controllers

import (
	"net/http"
	"news-go/src/dtos/article_dtos"
	"news-go/src/dtos/category_dtos"
	"news-go/src/helpers"
	"news-go/src/services/article_services"

	"github.com/gin-gonic/gin"
)

func GetAllArticles(service article_services.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pagination := helpers.GetPaginationParams(ctx)

		articles, total, err := service.GetPaginatedArticles(ctx, pagination)

		if err != nil {
			helpers.ErrorResponse(ctx, err, http.StatusInternalServerError)
			return
		}

		var articleDTOs []article_dtos.ArticleDTO
		for _, article := range articles {
			var categoryDTOs []category_dtos.CategoryDTO
			for _, category := range article.Categories {
				categoryDTOs = append(categoryDTOs, category_dtos.CategoryDTO{
					ID:    category.ID,
					UUID:  category.UUID,
					Label: category.Label,
					Value: category.Value,
				})
			}
			articleDTOs = append(articleDTOs, article_dtos.ArticleDTO{
				ID:          article.ID,
				UUID:        article.UUID,
				Title:       article.Title,
				Description: article.Description,
				URL:         article.URL,
				ImageURL:    article.ImageURL,
				Categories:  categoryDTOs,
				PublishedAt: article.PublishedAt,
				Source:      article.Source,
				CreatedAt:   article.CreatedAt,
				UpdatedAt:   article.UpdatedAt,
			})
		}

		helpers.SuccessResponse(ctx, "Data found!", articleDTOs, helpers.PaginationMeta{
			Page:  pagination.Page,
			Limit: pagination.Limit,
			Total: total,
		})

	}
}

func FetchArticles(service article_services.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := service.FetchAndSaveArticles(ctx); err != nil {
			helpers.ErrorResponse(ctx, err, http.StatusBadRequest)
			return
		}

		helpers.SuccessResponse(ctx, "article saved", nil)
	}

}
