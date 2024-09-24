package usecase

import (
	"context"
	"errors"

	dto "github.com/affrianr/go-mygram/dto/dto_photo"
	"github.com/affrianr/go-mygram/entity"
)

func (uc *PhotoUsecase) CreatePhoto(ctx context.Context, req dto.AddPhotoRequest) (entity.Photo, error) {

	photo := entity.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   req.UserID,
	}

	createdphoto, err := uc.PhotoRepository.CreatePhoto(ctx, photo)
	if err != nil {
		return entity.Photo{}, err
	}

	return createdphoto, nil
}

func (uc *PhotoUsecase) GetAllPhoto(ctx context.Context) ([]entity.Photo, error) {

	photo, err := uc.PhotoRepository.GetAllPhoto(ctx)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (uc *PhotoUsecase) GetPhotosByUser(ctx context.Context, userID uint) ([]entity.Photo, error) {

	photo, err := uc.PhotoRepository.GetPhotosByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (uc *PhotoUsecase) GetPhotoById(ctx context.Context, id uint) (entity.Photo, error) {
	photo, err := uc.PhotoRepository.GetPhotoById(ctx, id)
	if err != nil {
		return entity.Photo{}, err
	}

	if photo.ID == 0 {
		return entity.Photo{}, errors.New("social media not found")
	}

	return photo, nil
}

func (uc *PhotoUsecase) UpdatePhoto(ctx context.Context, req dto.UpdatePhotoRequest) (entity.Photo, error) {

	existingphoto, err := uc.PhotoRepository.GetPhotoById(ctx, req.ID)
	if err != nil {
		return entity.Photo{}, err
	}

	if existingphoto.ID == 0 {
		return entity.Photo{}, errors.New("social media not found")
	}

	updatedphoto, err := uc.PhotoRepository.UpdatePhoto(ctx, req)
	if err != nil {
		return entity.Photo{}, err
	}

	return updatedphoto, nil
}

func (uc *PhotoUsecase) DeletePhoto(ctx context.Context, id uint) error {
	existingphoto, err := uc.PhotoRepository.GetPhotoById(ctx, id)
	if err != nil {
		return err
	}

	if existingphoto.ID == 0 {
		return errors.New("social media not found")
	}

	err = uc.PhotoRepository.DeletePhoto(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
