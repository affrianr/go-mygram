package handler

import (
	"fmt"
	"net/http"
	"strconv"

	dto "github.com/affrianr/go-mygram/dto/dto_socialmedia"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var req dto.CreateSocialMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	req.UserID = uint(userData["id"].(float64))

	result, err := h.SocialMediaUsecase.CreateSocialMedia(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *SocialMediaHandler) GetAllSocialMedia(c *gin.Context) {

	result, err := h.SocialMediaUsecase.GetAllSocialMedia(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SocialMediaHandler) GetSocialMedia(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := h.SocialMediaUsecase.GetSocialMediaByUser(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SocialMediaHandler) GetSocialMediaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := h.SocialMediaUsecase.GetSocialMediaByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	var req dto.UpdateSocialMediaRequest
	fmt.Println(req, "usecase")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	req.ID = uint(id)

	result, err := h.SocialMediaUsecase.UpdateSocialMedia(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.SocialMediaUsecase.DeleteSocialMedia(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
