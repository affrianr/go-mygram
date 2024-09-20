package usecase

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_user"
	"github.com/affrianr/go-mygram/entity"
	errorcontext "github.com/affrianr/go-mygram/helpers"
	userRepo "github.com/affrianr/go-mygram/repository"
	"gorm.io/gorm"
)

type UserUsecase struct {
	UserRepository userRepo.IUserRepository
	GormDB         *gorm.DB
}

func InitUserUsecase(gormDB *gorm.DB, userRepository userRepo.IUserRepository) IUserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		GormDB:         gormDB}
}

type IUserUsecase interface {
	StoreUser(ctx context.Context, req dto.UserRegisterRequest) (user entity.User, err errorcontext.ErrorContext)
}
