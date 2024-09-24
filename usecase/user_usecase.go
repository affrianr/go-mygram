package usecase

import (
	"context"
	"errors"
	"fmt"

	dto "github.com/affrianr/go-mygram/dto/dto_user"
	"github.com/affrianr/go-mygram/entity"
	"github.com/affrianr/go-mygram/helpers"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) StoreUser(ctx context.Context, req dto.UserRegisterRequest) (user entity.User, err error) {
	// Hash password
	pwHash, bcryptErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if bcryptErr != nil {
		return entity.User{}, helpers.ErrHashPasswordFailed
	}

	user, err = u.UserRepository.GetUserByEmail(ctx, req.Email)

	if err != nil && errors.Is(err, helpers.ErrDataNotFound) {

		copier.Copy(&user, &req)
		fmt.Println(user)
		user.Password = string(pwHash)

		var userID uint
		userID, err = u.UserRepository.Create(ctx, user)
		if err != nil {
			return user, helpers.ErrUserStoreFailed
		}
		user.ID = userID
		fmt.Println(user)
		return user, nil
	} else if err != nil {
		return user, helpers.ErrUserCheckByEmailFailed
	}
	err = helpers.ErrDuplicateEmail
	return user, err
}

func (u *UserUsecase) Login(ctx context.Context, req dto.UserLoginRequest) (res dto.LoginResponse, err error) {
	findUser, err := u.UserRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, helpers.ErrDataNotFound) {
			return res, helpers.ErrInvalidEmailOrPassword
		}
		return res, helpers.ErrInternalServerError
	}
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(req.Password))
	if bcryptErr != nil {
		return res, helpers.ErrInvalidEmailOrPassword
	}

	token := helpers.GenerateToken(findUser.ID, findUser.Email)
	copier.Copy(&res.User, &findUser)
	resp := res
	resp.Token = token

	return resp, nil
}
