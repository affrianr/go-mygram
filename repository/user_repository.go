package repository

import (
	"context"

	"github.com/affrianr/go-mygram/entity"
	"github.com/affrianr/go-mygram/helpers"
)

func (r *UserRepository) Create(ctx context.Context, req entity.User) (userID uint, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	INSERT INTO public.users (username, email, password, age) VALUES(?, ?, ?, ?) RETURNING id
	`, req.Username, req.Email, req.Password, req.Age).Scan(&userID)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return userID, err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (user entity.User, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw("SELECT id, username, email, password, age, created_at, updated_at FROM public.users WHERE email = ?", email).Scan(&user)
	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}
	return user, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (user entity.User, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw("SELECT  id, username, email, age, created_at, updated_at FROM public.users WHERE id = ?", id).Scan(&user)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return user, err
}
