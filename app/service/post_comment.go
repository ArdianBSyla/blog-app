package service

import (
	"context"
	"errors"

	"github.com/personal/blog-app/app/dto"
	"github.com/personal/blog-app/app/model"
	"gorm.io/gorm"
)

type PostCommentService interface {
	CreatePostComment(ctx context.Context, createComment *dto.PostCommentToCreate) (*dto.PostComment, error)
	DeletePostComment(ctx context.Context, commentID int) error
	GetDeletedCommentsStats(ctx context.Context) (*dto.DeletedCommentsStats, error)
}

type postCommentService struct {
	orm *gorm.DB
}

func NewPostCommentService(orm *gorm.DB) PostCommentService {
	return &postCommentService{orm: orm}
}

func (s *postCommentService) CreatePostComment(ctx context.Context, createComment *dto.PostCommentToCreate) (*dto.PostComment, error) {
	comment := model.PostComment{
		PostID:  &createComment.PostID,
		Content: &createComment.Content,
		Author:  &createComment.Author,
	}

	err := s.orm.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return nil, err
	}

	return &dto.PostComment{
		ID:        comment.ID,
		CreatedAt: *comment.CreatedAt,
		UpdatedAt: *comment.UpdatedAt,
		PostCommentToCreate: dto.PostCommentToCreate{
			PostID:  *comment.PostID,
			Content: *comment.Content,
			Author:  *comment.Author,
		},
	}, nil
}

func (s *postCommentService) DeletePostComment(ctx context.Context, commentID int) error {
	var deletedPostComment model.PostComment

	result := s.orm.Where("id = ?", commentID).First(&deletedPostComment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("comment not found")
		}

		return result.Error
	}

	isDeleted := true
	deletedPostComment.IsDeleted = &isDeleted

	result = s.orm.Save(&deletedPostComment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *postCommentService) GetDeletedCommentsStats(ctx context.Context) (*dto.DeletedCommentsStats, error) {
	var stats dto.DeletedCommentsStats

	result := s.orm.Model(&model.PostComment{}).
		Select("SUM(CASE WHEN is_deleted THEN 1 ELSE 0 END) as deleted_count, COUNT(*) as total_count").
		Scan(&stats)

	if result.Error != nil {
		return nil, result.Error
	}

	return &stats, nil
}
