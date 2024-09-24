package usecase

import (
	"context"
	"errors"

	dto "github.com/affrianr/go-mygram/dto/dto_comment"
	"github.com/affrianr/go-mygram/entity"
)

func (uc *CommentUsecase) CreateComment(ctx context.Context, req dto.AddCommentRequest) (entity.Comment, error) {

	Comment := entity.Comment{
		UserID:  req.UserID,
		PhotoID: req.PhotoID,
		Message: req.Message,
	}

	createdComment, err := uc.CommentRepository.CreateComment(ctx, Comment)
	if err != nil {
		return entity.Comment{}, err
	}

	return createdComment, nil
}

func (uc *CommentUsecase) GetAllComment(ctx context.Context) ([]entity.Comment, error) {

	Comment, err := uc.CommentRepository.GetAllComment(ctx)
	if err != nil {
		return nil, err
	}

	return Comment, nil
}

func (uc *CommentUsecase) GetCommentsByUser(ctx context.Context, userID uint) ([]entity.Comment, error) {

	Comment, err := uc.CommentRepository.GetCommentByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return Comment, nil
}

func (uc *CommentUsecase) GetCommentById(ctx context.Context, id uint) (entity.Comment, error) {
	Comment, err := uc.CommentRepository.GetCommentById(ctx, id)
	if err != nil {
		return entity.Comment{}, err
	}

	if Comment.ID == 0 {
		return entity.Comment{}, errors.New("social media not found")
	}

	return Comment, nil
}

func (uc *CommentUsecase) UpdateComment(ctx context.Context, req dto.UpdateCommentRequest) (entity.Comment, error) {

	existingComment, err := uc.CommentRepository.GetCommentById(ctx, req.ID)
	if err != nil {
		return entity.Comment{}, err
	}

	if existingComment.ID == 0 {
		return entity.Comment{}, errors.New("social media not found")
	}

	updatedComment, err := uc.CommentRepository.UpdateComment(ctx, req)
	if err != nil {
		return entity.Comment{}, err
	}

	return updatedComment, nil
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, id uint) error {
	existingComment, err := uc.CommentRepository.GetCommentById(ctx, id)
	if err != nil {
		return err
	}

	if existingComment.ID == 0 {
		return errors.New("social media not found")
	}

	err = uc.CommentRepository.DeleteComment(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
