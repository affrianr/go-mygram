package dto

type AddSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
	UserID         uint   `json:"user_id" binding:"required"`
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
	UserID         uint   `json:"user_id"`
}

type UpdateSocialMediaRequest struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}
