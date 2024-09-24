package repository

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_comment"
	"github.com/affrianr/go-mygram/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	gormDB *gorm.DB
}

type ICommentRepository interface {
	CreateComment(ctx context.Context, req entity.Comment) (comment entity.Comment, err error)
	GetAllComment(ctx context.Context) (comments []entity.Comment, err error)
	GetCommentByUser(ctx context.Context, userID uint) (comments []entity.Comment, err error)
	GetCommentById(ctx context.Context, id uint) (comment entity.Comment, err error)
	UpdateComment(ctx context.Context, req dto.UpdateCommentRequest) (comment entity.Comment, err error)
	DeleteComment(ctx context.Context, id uint) error
}

func InitCommentRepository(gormDB *gorm.DB) ICommentRepository {
	return &CommentRepository{gormDB: gormDB}
}
