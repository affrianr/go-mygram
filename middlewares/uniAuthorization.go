package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/affrianr/go-mygram/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UnifiedAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.DB
		fullPath := c.FullPath()
		pathParts := strings.Split(fullPath, "/")

		var (
			paramName string
			tableName string
		)

		if len(pathParts) < 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "invalid path",
			})
			return
		}

		switch pathParts[1] {
		case "users":
			paramName = "id"
			tableName = "users"
		case "photos":
			paramName = "id"
			tableName = "photos"
		case "comments":
			paramName = "id"
			tableName = "comments"
		case "social-media":
			paramName = "id"
			tableName = "social_media"
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "invalid path",
			})
			return
		}

		// Extract the ID from the path parameters
		paramValue := c.Param(paramName)
		id, err := strconv.Atoi(paramValue)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "invalid param",
			})
			return
		}

		// Get the user ID from the JWT token
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		// Prepare the SQL query
		var query string
		var args []interface{}

		if tableName == "users" {
			query = "SELECT id FROM users WHERE id = ?"
			args = []interface{}{id}
		} else {
			query = fmt.Sprintf("SELECT user_id FROM %s WHERE id = ?", tableName)
			args = []interface{}{id}
		}

		// Execute the raw SQL query
		var result struct {
			UserID uint `gorm:"column:user_id"`
		}
		err = db.Raw(query, args...).Scan(&result).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Not found",
					"message": "data does not exist",
				})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal server error",
					"message": "failed to query database",
				})
			}
			return
		}

		// For users table, we compare with the ID directly
		var entityUserID uint
		if tableName == "users" {
			entityUserID = result.UserID
		} else {
			entityUserID = result.UserID
		}

		// Check if the user is authorized
		if entityUserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not authorized to access this data",
			})
			return
		}

		c.Next()
	}
}
