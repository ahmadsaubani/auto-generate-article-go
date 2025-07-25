package articles

import (
	"news-go/src/entities/categories"
	"time"
)

type Article struct {
	ID            uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID          string   `gorm:"type:uuid;uniqueIndex" json:"uuid"`
	Title         string   `json:"title" db:"title"`
	Description   string   `json:"description" db:"description"`
	URL           string   `json:"url" db:"url"`
	ImageURL      string   `json:"image_url" db:"image_url"`
	CategoryNames []string `gorm:"-" json:"category_names"` // hanya untuk proses, tidak masuk DB

	Categories  []categories.Category `gorm:"many2many:article_categories;" json:"categories"`
	PublishedAt time.Time             `json:"published_at" db:"published_at"`
	Source      *string               `json:"source" db:"source"`
	CreatedAt   time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
}
