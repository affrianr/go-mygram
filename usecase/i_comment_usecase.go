package usecase

import (
	"context"

	dto "github.com/affrianr/go-mygram/dto/dto_comment"
	"github.com/affrianr/go-mygram/entity"
	commentRepo "github.com/affrianr/go-mygram/repository"
	"gorm.io/gorm"
)

type CommentUsecase struct {
	CommentRepository commentRepo.ICommentRepository
	GormDB            *gorm.DB
}

// InitSocialMediaUsecase initializes the SocialMediaUsecase and returns the use case interface
func InitCommentUsecase(gormDB *gorm.DB, commentRepository commentRepo.ICommentRepository) ICommentUsecase {
	return &CommentUsecase{
		CommentRepository: commentRepository,
		GormDB:            gormDB,
	}
}

type ICommentUsecase interface {
	CreateComment(ctx context.Context, req dto.AddCommentRequest) (entity.Comment, error)
	GetAllComment(ctx context.Context) ([]entity.Comment, error)
	GetCommentsByUser(ctx context.Context, userID uint) ([]entity.Comment, error)
	GetCommentById(ctx context.Context, id uint) (entity.Comment, error)
	UpdateComment(ctx context.Context, req dto.UpdateCommentRequest) (entity.Comment, error)
	DeleteComment(ctx context.Context, id uint) error
}
