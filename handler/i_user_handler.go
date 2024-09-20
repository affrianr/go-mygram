package handler

import userUsecase "github.com/affrianr/go-mygram/usecase"

type UserHandler struct {
	UserCase userUsecase.IUserUsecase
}

func InitUserHandler(userUsecase userUsecase.IUserUsecase) *UserHandler {
	return &UserHandler{UserCase: userUsecase}
}
