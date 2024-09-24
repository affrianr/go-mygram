package repository

import (
	"context"

	"github.com/affrianr/go-mygram/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	gormDB *gorm.DB
}

type IUserRepository interface {
	Create(ctx context.Context, req entity.User) (userID uint, err error)
	GetUserByEmail(ctx context.Context, email string) (user entity.User, err error)
	GetUserByID(ctx context.Context, id uint) (user entity.User, err error)
}

func InitUserRepository(gormDB *gorm.DB) IUserRepository {
	return &UserRepository{gormDB: gormDB}
}
