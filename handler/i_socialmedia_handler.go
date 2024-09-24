package handler

import "github.com/affrianr/go-mygram/usecase"

type SocialMediaHandler struct {
	SocialMediaUsecase usecase.ISocialMediaUsecase
}

func InitSocialMediaHandler(usecase usecase.ISocialMediaUsecase) *SocialMediaHandler {
	return &SocialMediaHandler{
		SocialMediaUsecase: usecase,
	}
}
