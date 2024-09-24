package repository

import (
	"context"
	"fmt"

	dto "github.com/affrianr/go-mygram/dto/dto_socialmedia"
	"github.com/affrianr/go-mygram/entity"
	"github.com/affrianr/go-mygram/helpers"
)

func (r *SocialMediaRepository) CreateSocialMedia(ctx context.Context, req entity.SocialMedia) (socmed entity.SocialMedia, err error) {
	db := r.gormDB.WithContext(ctx)
	fmt.Println(req)

	result := db.Raw(`
	INSERT INTO public.social_media (name, social_media_url, user_id) VALUES(?, ?, ?) RETURNING *
	`, req.Name, req.SocialMediaUrl, req.UserID).Scan(&socmed)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return socmed, err
}

func (r *SocialMediaRepository) GetAllSocialMedia(ctx context.Context) (socmed []entity.SocialMedia, err error) {
	db := r.gormDB.WithContext(ctx)

	socmeds := db.Raw(`
	SELECT id, name, social_media_url, user_id, created_at, updated_at FROM public.social_media 
	`).Scan(&socmed)

	if socmeds.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if socmeds.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return socmed, err
}

func (r *SocialMediaRepository) GetSocialMediaByUser(ctx context.Context, userID uint) (socmed []entity.SocialMedia, err error) {
	db := r.gormDB.WithContext(ctx)

	socmeds := db.Raw(`
	SELECT id, name, social_media_url, user_id, created_at, updated_at  FROM public.social_media WHERE user_id = ?
	`, userID).Scan(&socmed)

	if socmeds.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if socmeds.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return socmed, err
}

func (r *SocialMediaRepository) GetSocialMediaById(ctx context.Context, id uint) (socmed entity.SocialMedia, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, name, social_media_url, user_id, created_at, updated_at FROM public.social_media WHERE id = ?
	`, id).Scan(&socmed)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return socmed, err
}

func (r *SocialMediaRepository) UpdateSocialMedia(ctx context.Context, req dto.UpdateSocialMediaRequest) (socmed entity.SocialMedia, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	UPDATE public.social_media 
	SET name = ?, social_media_url = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	RETURNING *
	`, req.Name, req.SocialMediaUrl, req.ID).Scan(&socmed)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return socmed, err
}

func (r *SocialMediaRepository) DeleteSocialMedia(ctx context.Context, id uint) error {
	db := r.gormDB.WithContext(ctx)

	result := db.Exec(`
	DELETE FROM public.social_media WHERE id = ?
	`, id)

	if result.Error != nil {
		return helpers.ErrInternalServerError
	} else if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}

	return nil
}
