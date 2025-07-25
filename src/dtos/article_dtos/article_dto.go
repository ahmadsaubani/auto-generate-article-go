package article_dtos

import (
	"news-go/src/dtos/category_dtos"
	"time"
)

type ArticleDTO struct {
	ID          uint   `json:"id"`
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`

	Categories  []category_dtos.CategoryDTO `json:"categories"`
	PublishedAt time.Time                   `json:"published_at"`
	Source      *string                     `json:"source"`
	CreatedAt   time.Time                   `json:"created_at"`
	UpdatedAt   time.Time                   `json:"updated_at"`
}
