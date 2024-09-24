package routers

import (
	"github.com/affrianr/go-mygram/handler"
	"github.com/affrianr/go-mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCommentRoutes(Route *gin.Engine, commentHandler *handler.CommentHandler) *gin.Engine {

	comment := Route.Group("/Comments")
	{
		comment.GET("/user/:user_id", commentHandler.GetCommentsByUser)
		comment.GET("/", commentHandler.GetAllComment)
		comment.Use(middlewares.Authentication())
		comment.POST("/", commentHandler.CreateComment)
		comment.GET("/:id", commentHandler.GetCommentById)
		comment.PUT("/:id", middlewares.UnifiedAuthorization(), commentHandler.UpdateComment)
		comment.DELETE("/:id", middlewares.UnifiedAuthorization(), commentHandler.DeleteComment)
	}

	return Route
}
