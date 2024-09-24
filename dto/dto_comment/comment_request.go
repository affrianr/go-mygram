package dto

type AddCommentRequest struct {
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

type UpdateCommentRequest struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}
