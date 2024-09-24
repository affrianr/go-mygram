package routers

import (
	"github.com/affrianr/go-mygram/handler"
	"github.com/affrianr/go-mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSocmedRoutes(Route *gin.Engine, socialMediaHandler *handler.SocialMediaHandler) *gin.Engine {

	socialMedia := Route.Group("/social-media")
	{
		socialMedia.GET("/user/:user_id", socialMediaHandler.GetSocialMedia)
		socialMedia.GET("/", socialMediaHandler.GetAllSocialMedia)
		socialMedia.Use(middlewares.Authentication())
		socialMedia.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMedia.GET("/:id", socialMediaHandler.GetSocialMediaByID)
		socialMedia.PUT("/:id", middlewares.UnifiedAuthorization(), socialMediaHandler.UpdateSocialMedia)
		socialMedia.DELETE("/:id", middlewares.UnifiedAuthorization(), socialMediaHandler.DeleteSocialMedia)
	}

	return Route
}
