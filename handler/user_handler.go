package handler

import (
	"net/http"

	dto "github.com/affrianr/go-mygram/dto/dto_user"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.UserRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.UserCase.StoreUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.UserRegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      uint(user.Age),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully registered",
		"user":    res,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.UserCase.Login(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "invalid email or password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    res.User,
		"token":   res.Token,
	})
}
