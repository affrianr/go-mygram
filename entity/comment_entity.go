package entity

import (
	"time"
)

type Comment struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id" gorm:"not null"`
	User      User
	PhotoID   uint `json:"photo_id" gorm:"not null"`
	Photo     Photo
	Message   string    `json:"message" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
