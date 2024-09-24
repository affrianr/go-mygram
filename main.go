package main

import (
	"os"

	"github.com/affrianr/go-mygram/config"
	"github.com/affrianr/go-mygram/handler"
	"github.com/affrianr/go-mygram/repository"
	routers "github.com/affrianr/go-mygram/routes"
	"github.com/affrianr/go-mygram/usecase"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDB()
}

func main() {

	Routes := gin.Default()
	userRepository := repository.InitUserRepository(config.DB)
	userUsecase := usecase.InitUserUsecase(config.DB, userRepository)
	userHandler := handler.InitUserHandler(userUsecase)
	routers.InitUserRoutes(Routes, userHandler)

	socmedRepository := repository.InitSocialMediaRepository(config.DB)
	socmedUsecase := usecase.InitSocialMediaUsecase(config.DB, socmedRepository)
	socmedHandler := handler.InitSocialMediaHandler(socmedUsecase)
	routers.InitSocmedRoutes(Routes, socmedHandler)

	photoRepository := repository.InitPhotoRepository(config.DB)
	photoUsecase := usecase.InitPhotoUsecase(config.DB, photoRepository)
	photoHandler := handler.IniPhotoHandler(photoUsecase)
	routers.InitPhotoRoutes(Routes, photoHandler)

	commentRepository := repository.InitCommentRepository(config.DB)
	commentUsecase := usecase.InitCommentUsecase(config.DB, commentRepository)
	commentHandler := handler.IniCommentHandler(commentUsecase)
	routers.InitCommentRoutes(Routes, commentHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080" // Default port
	}

	Routes.Run(port)
}
