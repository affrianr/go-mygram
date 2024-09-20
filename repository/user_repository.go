package repository

import (
	"context"
	"net/http"

	"github.com/affrianr/go-mygram/entity"
	errorcontext "github.com/affrianr/go-mygram/helpers"
)

func (r *UserRepository) Register(ctx context.Context, req entity.User) (userID string, err errorcontext.ErrorContext) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	INSERT INTO public.users (username, email, password, age) VALUE(?, ?, ?, ?) RETURNING id
	`, req.Username, req.Email, req.Password, req.Age).Scan(&userID)

	if result.Error != nil {
		err = errorcontext.NewAppError(result.Error, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if result.RowsAffected == 0 {
		err = errorcontext.NewAppError(nil, http.StatusNotFound, "Data not found")
	}

	return
}
