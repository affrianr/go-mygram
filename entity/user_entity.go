package entity

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"uniqueIndex;not null"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null"`
	Password    string    `json:"password" gorm:"not null;check:length(password) >= 6"`
	Age         int       `json:"age" gorm:"check:age>8"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
