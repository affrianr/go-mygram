package repository

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_comment"
	"github.com/affrianr/go-mygram/entity"
	"github.com/affrianr/go-mygram/helpers"
)

func (r *CommentRepository) CreateComment(ctx context.Context, req entity.Comment) (comment entity.Comment, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	INSERT INTO public.comments (user_id, photo_id, message) VALUES(?, ?, ?) RETURNING *
	`, req.UserID, req.PhotoID, req.Message).Scan(&comment)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return comment, err
}

func (r *CommentRepository) GetAllComment(ctx context.Context) (comments []entity.Comment, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, user_id, photo_id, message, created_at, updated_at FROM public.comments
	`).Scan(&comments)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return comments, err
}

func (r *CommentRepository) GetCommentByUser(ctx context.Context, userID uint) (comments []entity.Comment, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, user_id, photo_id, message, created_at, updated_at FROM public.comments WHERE id = ?
	`, userID).Scan(&comments)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return comments, err
}

func (r *CommentRepository) GetCommentById(ctx context.Context, id uint) (comment entity.Comment, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	SELECT id, user_id, photo_id, message, created_at, updated_at FROM public.comments WHERE id = ?
	`, id).Scan(&comment)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return comment, err
}

func (r *CommentRepository) UpdateComment(ctx context.Context, req dto.UpdateCommentRequest) (comment entity.Comment, err error) {
	db := r.gormDB.WithContext(ctx)

	result := db.Raw(`
	UPDATE public.comments
	SET message = ? updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	RETURNING *
	`, req.Message, req.ID).Scan(&comment)

	if result.Error != nil {
		err = helpers.ErrInternalServerError
		return
	} else if result.RowsAffected == 0 {
		err = helpers.ErrDataNotFound
	}

	return comment, err
}

func (r *CommentRepository) DeleteComment(ctx context.Context, id uint) error {
	db := r.gormDB.WithContext(ctx)

	result := db.Exec(`
	DELETE FROM public.comments WHERE id = ?
	`, id)

	if result.Error != nil {
		return helpers.ErrInternalServerError
	} else if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}

	return nil
}
