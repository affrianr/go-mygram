package routers

import (
	"github.com/affrianr/go-mygram/handler"
	"github.com/affrianr/go-mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func InitPhotoRoutes(Route *gin.Engine, photoHandler *handler.PhotoHandler) *gin.Engine {

	photo := Route.Group("/photos")
	{
		photo.GET("/user/:user_id", photoHandler.GetPhotosByUser)
		photo.GET("/", photoHandler.GetAllPhoto)
		photo.Use(middlewares.Authentication())
		photo.POST("/", photoHandler.CreatePhoto)
		photo.GET("/:id", photoHandler.GetPhotoById)
		photo.PUT("/:id", middlewares.UnifiedAuthorization(), photoHandler.UpdatePhoto)
		photo.DELETE("/:id", middlewares.UnifiedAuthorization(), photoHandler.DeletePhoto)
	}

	return Route
}
