package routers

import (
	"github.com/affrianr/go-mygram/handler"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(Routes *gin.Engine, handler *handler.UserHandler) {
	userRouter := Routes.Group("/users")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}
}
