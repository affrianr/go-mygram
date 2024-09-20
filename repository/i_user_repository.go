package repository

import (
	"context"

	"github.com/affrianr/go-mygram/entity"
	errorcontext "github.com/affrianr/go-mygram/helpers"
	"gorm.io/gorm"
)

type UserRepository struct {
	gormDB *gorm.DB
}

type IUserRepository interface {
	Register(ctx context.Context, req entity.User) (userID string, err errorcontext.ErrorContext)
	// Login(ctx context.Context, req entity.UserLogin) (userID string, err errorcontext.ErrorContext)
}

func InitUserRepository(gormDB *gorm.DB) IUserRepository {
	return &UserRepository{gormDB: gormDB}
}
