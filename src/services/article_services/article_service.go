package article_services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"news-go/src/entities/articles"
	"news-go/src/helpers"
	"news-go/src/repositories/article_repositories"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleResponse struct {
	Data []struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		URL         string   `json:"url"`
		ImageURL    string   `json:"image_url"`
		PublishedAt string   `json:"published_at"`
		Categories  []string `json:"categories"`
		Source      *string  `json:"source"`
	} `json:"data"`
}
type ArticleService interface {
	FetchAndSaveArticles(ctx *gin.Context) error
	GetAll(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, error)
	GetPaginatedArticles(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, int64, error)
}

type articleService struct {
	repo article_repositories.ArticleRepository
}

func NewArticleService(repo article_repositories.ArticleRepository) ArticleService {
	return &articleService{repo}
}

func (s *articleService) GetPaginatedArticles(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, int64, error) {
	articleList, err := s.repo.GetAll(ctx, pag)
	if err != nil {
		return nil, 0, fmt.Errorf("sorry, we encountered an issue fetching the article list. Please try again later: %w", err)
	}

	total, err := helpers.CountModel[articles.Article]()
	if err != nil {
		return nil, 0, fmt.Errorf("sorry, we couldn't count the article at the moment. Please try again later: %w", err)
	}

	return articleList, total, nil
}

func (s *articleService) GetAll(ctx *gin.Context, pag helpers.PaginationParams) ([]articles.Article, error) {
	articleLists, err := s.repo.GetAll(ctx, pag)
	if err != nil {
		return nil, fmt.Errorf("sorry, we encountered an issue fetching articles. Please try again later: %w", err)
	}

	return articleLists, nil
}

func (s *articleService) FetchAndSaveArticles(ctx *gin.Context) error {
	baseURL, _ := url.Parse("https://api.thenewsapi.com/v1/news/all")
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		log.Fatal("API_TOKEN is missing in environment")
	}
	params := url.Values{}
	params.Add("api_token", apiToken)
	params.Add("categories", "general,science,business,tech,sports,travel")
	params.Add("language", "en")

	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var news ArticleResponse
	if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
		return err
	}

	var articlesList []articles.Article
	for _, item := range news.Data {
		pubTime, _ := time.Parse(time.RFC3339, item.PublishedAt)

		articlesList = append(articlesList, articles.Article{
			Title:         item.Title,
			Description:   item.Description,
			URL:           item.URL,
			ImageURL:      item.ImageURL,
			CategoryNames: item.Categories,
			PublishedAt:   pubTime,
			Source:        item.Source,
		})
	}

	return s.repo.SaveArticles(ctx, articlesList)
}
