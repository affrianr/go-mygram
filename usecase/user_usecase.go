package usecase

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_user"
	"github.com/affrianr/go-mygram/entity"
	errorcontext "github.com/affrianr/go-mygram/helpers"
	"github.com/jinzhu/copier"
)

func (u *UserUsecase) StoreUser(ctx context.Context, req dto.UserRegisterRequest) (user entity.User, err errorcontext.ErrorContext) {

	copier.Copy(&user, &req)
	_, err = u.StoreUser(ctx, req)
	if err != nil {
		return
	}
	return user, err
}
