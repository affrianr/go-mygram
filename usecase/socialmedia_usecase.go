package usecase

import (
	"context"
	"errors"

	dto "github.com/affrianr/go-mygram/dto/dto_socialmedia"
	"github.com/affrianr/go-mygram/entity"
)

func (uc *SocialMediaUsecase) CreateSocialMedia(ctx context.Context, req dto.CreateSocialMediaRequest) (entity.SocialMedia, error) {

	socmed := entity.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         req.UserID,
	}
	createdSocmed, err := uc.SocialMediaRepository.CreateSocialMedia(ctx, socmed)
	if err != nil {
		return entity.SocialMedia{}, err
	}

	return createdSocmed, nil
}

func (uc *SocialMediaUsecase) GetAllSocialMedia(ctx context.Context) ([]entity.SocialMedia, error) {
	socmed, err := uc.SocialMediaRepository.GetAllSocialMedia(ctx)
	if err != nil {
		return nil, err
	}

	return socmed, nil
}

func (uc *SocialMediaUsecase) GetSocialMediaByUser(ctx context.Context, userID uint) ([]entity.SocialMedia, error) {
	socmed, err := uc.SocialMediaRepository.GetSocialMediaByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return socmed, nil
}

func (uc *SocialMediaUsecase) GetSocialMediaByID(ctx context.Context, id uint) (entity.SocialMedia, error) {
	socmed, err := uc.SocialMediaRepository.GetSocialMediaById(ctx, id)
	if err != nil {
		return entity.SocialMedia{}, err
	}

	if socmed.ID == 0 {
		return entity.SocialMedia{}, errors.New("social media not found")
	}

	return socmed, nil
}

func (uc *SocialMediaUsecase) UpdateSocialMedia(ctx context.Context, req dto.UpdateSocialMediaRequest) (entity.SocialMedia, error) {
	existingSocmed, err := uc.SocialMediaRepository.GetSocialMediaById(ctx, req.ID)
	if err != nil {
		return entity.SocialMedia{}, err
	}

	if existingSocmed.ID == 0 {
		return entity.SocialMedia{}, errors.New("social media not found")
	}

	updatedSocmed, err := uc.SocialMediaRepository.UpdateSocialMedia(ctx, req)
	if err != nil {
		return entity.SocialMedia{}, err
	}

	return updatedSocmed, nil
}

func (uc *SocialMediaUsecase) DeleteSocialMedia(ctx context.Context, id uint) error {
	existingSocmed, err := uc.SocialMediaRepository.GetSocialMediaById(ctx, id)
	if err != nil {
		return err
	}

	if existingSocmed.ID == 0 {
		return errors.New("social media not found")
	}

	err = uc.SocialMediaRepository.DeleteSocialMedia(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
