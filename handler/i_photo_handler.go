package handler

import "github.com/affrianr/go-mygram/usecase"

type PhotoHandler struct {
	PhotoUsecase usecase.IPhotoUsecase
}

func IniPhotoHandler(usecase usecase.IPhotoUsecase) *PhotoHandler {
	return &PhotoHandler{
		PhotoUsecase: usecase,
	}
}
