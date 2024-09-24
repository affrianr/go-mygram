package entity

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         uint   `gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
