package repository

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_photo"
	"github.com/affrianr/go-mygram/entity"
	"github.com/affrianr/go-mygram/helpers"
)

func (r *PhotoRepository) CreatePhoto(ctx context.Context, req entity.Photo) (photo entity.Photo, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	INSERT INTO public.photos (title, caption, photo_url, user_id) VALUES(?, ?, ?, ?) RETURNING *
	`, req.Title, req.Caption, req.PhotoUrl, req.UserID).Scan(&photo)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return photo, err
}

func (r *PhotoRepository) GetAllPhoto(ctx context.Context) (photos []entity.Photo, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, title, caption, photo_url, user_id, created_at, updated_at FROM public.photos 
	`).Scan(&photos)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return photos, err
}

func (r *PhotoRepository) GetPhotosByUser(ctx context.Context, userID uint) (photo []entity.Photo, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, title, caption, photo_url, user_id, created_at, updated_at FROM public.photos WHERE user_id = ?
	`, userID).Scan(&photo)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return photo, err
}

func (r *PhotoRepository) GetPhotoById(ctx context.Context, id uint) (photo entity.Photo, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, title, caption, photo_url, user_id, created_at, updated_at FROM public.photos WHERE id = ?
	`, id).Scan(&photo)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return photo, err
}

func (r *PhotoRepository) UpdatePhoto(ctx context.Context, req dto.UpdatePhotoRequest) (photo entity.Photo, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	UPDATE public.photos
	SET title = ?, caption = ?, photo_url = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	RETURNING *
	`, req.Title, req.Caption, req.PhotoUrl, req.ID).Scan(&photo)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return photo, err
}

func (r *PhotoRepository) DeletePhoto(ctx context.Context, id uint) error {
	db := r.gormDB.WithContext(ctx)

	result := db.Exec(`
	DELETE FROM public.photos WHERE id = ?
	`, id)

	if result.Error != nil {
		return helpers.ErrInternalServerError
	} else if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}

	return nil
}
