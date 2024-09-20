package entity

import (
	"time"
)

type Photo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url" gorm:"not null"`
	UserID    uint   `json:"user_id" gorm:"not null"`
	User      User
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()`
	Comments  []Comment
}
