package categories

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UUID      string    `gorm:"type:uuid;uniqueIndex" json:"uuid"`
	Label     string    `gorm:"not null" json:"label"`
	Value     string    `gorm:"not null" json:"value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
