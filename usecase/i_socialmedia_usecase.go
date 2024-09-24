package usecase

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_socialmedia"
	"github.com/affrianr/go-mygram/entity"
	socmedRepo "github.com/affrianr/go-mygram/repository"
	"gorm.io/gorm"
)

type SocialMediaUsecase struct {
	SocialMediaRepository socmedRepo.ISocialMediaRepository
	GormDB                *gorm.DB
}

// InitSocialMediaUsecase initializes the SocialMediaUsecase and returns the use case interface
func InitSocialMediaUsecase(gormDB *gorm.DB, socialMediaRepository socmedRepo.ISocialMediaRepository) ISocialMediaUsecase {
	return &SocialMediaUsecase{
		SocialMediaRepository: socialMediaRepository,
		GormDB:                gormDB,
	}
}

// ISocialMediaUsecase defines the methods for the Social Media Usecase
type ISocialMediaUsecase interface {
	CreateSocialMedia(ctx context.Context, req dto.CreateSocialMediaRequest) (entity.SocialMedia, error)
	GetAllSocialMedia(ctx context.Context) ([]entity.SocialMedia, error)
	GetSocialMediaByUser(ctx context.Context, userID uint) ([]entity.SocialMedia, error)
	GetSocialMediaByID(ctx context.Context, id uint) (entity.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, req dto.UpdateSocialMediaRequest) (entity.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id uint) error
}
