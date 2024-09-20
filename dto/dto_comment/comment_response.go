package dto

import "time"

type AddCommentResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCommentResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
	User    struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
