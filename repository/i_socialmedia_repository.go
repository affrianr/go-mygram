package repository

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_socialmedia"
	"github.com/affrianr/go-mygram/entity"
	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	gormDB *gorm.DB
}

type ISocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, req entity.SocialMedia) (socmed entity.SocialMedia, err error)
	GetAllSocialMedia(ctx context.Context) (socmed []entity.SocialMedia, err error)
	GetSocialMediaByUser(ctx context.Context, userID uint) (socmed []entity.SocialMedia, err error)
	GetSocialMediaById(ctx context.Context, id uint) (socmed entity.SocialMedia, err error)
	UpdateSocialMedia(ctx context.Context, req dto.UpdateSocialMediaRequest) (socmed entity.SocialMedia, err error)
	DeleteSocialMedia(ctx context.Context, id uint) error
}

func InitSocialMediaRepository(gormDB *gorm.DB) ISocialMediaRepository {
	return &SocialMediaRepository{gormDB: gormDB}
}
