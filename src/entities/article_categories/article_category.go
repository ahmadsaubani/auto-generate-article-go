package article_categories

type ArticleCategory struct {
	ArticleID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}
