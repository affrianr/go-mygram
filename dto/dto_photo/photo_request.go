package dto

type AddPhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
}
