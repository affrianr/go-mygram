package repository

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_photo"
	"github.com/affrianr/go-mygram/entity"
	"gorm.io/gorm"
)

type PhotoRepository struct {
	gormDB *gorm.DB
}

type IPhotoRepository interface {
	CreatePhoto(ctx context.Context, req entity.Photo) (photo entity.Photo, err error)
	GetAllPhoto(ctx context.Context) (photos []entity.Photo, err error)
	GetPhotosByUser(ctx context.Context, userID uint) (photo []entity.Photo, err error)
	GetPhotoById(ctx context.Context, id uint) (photo entity.Photo, err error)
	UpdatePhoto(ctx context.Context, req dto.UpdatePhotoRequest) (photo entity.Photo, err error)
	DeletePhoto(ctx context.Context, id uint) error
}

func InitPhotoRepository(gormDB *gorm.DB) IPhotoRepository {
	return &PhotoRepository{gormDB: gormDB}
}
