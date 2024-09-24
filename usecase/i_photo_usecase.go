package usecase

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_photo"
	"github.com/affrianr/go-mygram/entity"
	photoRepo "github.com/affrianr/go-mygram/repository"
	"gorm.io/gorm"
)

type PhotoUsecase struct {
	PhotoRepository photoRepo.IPhotoRepository
	GormDB          *gorm.DB
}

// InitSocialMediaUsecase initializes the SocialMediaUsecase and returns the use case interface
func InitPhotoUsecase(gormDB *gorm.DB, photoRepository photoRepo.IPhotoRepository) IPhotoUsecase {
	return &PhotoUsecase{
		PhotoRepository: photoRepository,
		GormDB:          gormDB,
	}
}

// IPhotoUsecase defines the methods for the Social Media Usecase
type IPhotoUsecase interface {
	CreatePhoto(ctx context.Context, req dto.AddPhotoRequest) (entity.Photo, error)
	GetAllPhoto(ctx context.Context) ([]entity.Photo, error)
	GetPhotosByUser(ctx context.Context, userID uint) ([]entity.Photo, error)
	GetPhotoById(ctx context.Context, id uint) (entity.Photo, error)
	UpdatePhoto(ctx context.Context, req dto.UpdatePhotoRequest) (entity.Photo, error)
	DeletePhoto(ctx context.Context, id uint) error
}
