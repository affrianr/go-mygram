package handler

import (
	"net/http"

	dto "github.com/affrianr/go-mygram/dto/dto_user"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) StoreUser(ctx *gin.Context) {
	req := dto.UserRegisterRequest{}

	res, err := h.UserCase.StoreUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success created user",
		"data":    res,
	})
}
