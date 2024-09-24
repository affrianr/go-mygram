package handler

import "github.com/affrianr/go-mygram/usecase"

type CommentHandler struct {
	CommentUsecase usecase.ICommentUsecase
}

func IniCommentHandler(usecase usecase.ICommentUsecase) *CommentHandler {
	return &CommentHandler{
		CommentUsecase: usecase,
	}
}
