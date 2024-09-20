package dto

import "time"

type AddPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required"`
	UserID    uint      `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	User     struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
