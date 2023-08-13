package service

import (
	"context"
	"errors"

	"github.com/personal/blog-app/app/dto"
	"github.com/personal/blog-app/app/helper"
	"github.com/personal/blog-app/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogPostService interface {
	CreateBlogPost(ctx context.Context, createBlogPost *dto.BlogPostToCreate) (*dto.BlogPost, error)
	GetBlogPostByID(ctx context.Context, id int) (*dto.BlogPost, error)
	GetListOfBlogPosts(ctx context.Context, limit, offset int) ([]*dto.BlogPost, error)
	UpdateBlogPost(ctx context.Context, postId int, updateBlogPost *dto.UpdateBlogPost) (*dto.BlogPost, error)
	DeleteBlogPost(ctx context.Context, postId int) error
}

type blogPostService struct {
	orm *gorm.DB
}

func NewBlogPostService(orm *gorm.DB) BlogPostService {
	return &blogPostService{orm: orm}
}

func (s *blogPostService) CreateBlogPost(ctx context.Context, createBlogPost *dto.BlogPostToCreate) (*dto.BlogPost, error) {
	blogPost := model.BlogPost{
		Title:   &createBlogPost.Title,
		Content: &createBlogPost.Content,
		Author:  &createBlogPost.Author,
	}

	err := s.orm.WithContext(ctx).Create(&blogPost).Error
	if err != nil {
		return nil, err
	}

	return &dto.BlogPost{
		ID:        blogPost.ID,
		CreatedAt: *blogPost.CreatedAt,
		UpdatedAt: *blogPost.UpdatedAt,
		BlogPostToCreate: dto.BlogPostToCreate{
			Title:   *blogPost.Title,
			Content: *blogPost.Content,
			Author:  *blogPost.Author,
		},
	}, nil
}

func (s *blogPostService) GetBlogPostByID(ctx context.Context, id int) (*dto.BlogPost, error) {

	var post model.BlogPost

	result := s.orm.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog post not found")
		}

		return nil, result.Error
	}

	return &dto.BlogPost{
		ID:        post.ID,
		CreatedAt: *post.CreatedAt,
		UpdatedAt: *post.UpdatedAt,
		BlogPostToCreate: dto.BlogPostToCreate{
			Title:   *post.Title,
			Content: *post.Content,
			Author:  *post.Author,
		},
	}, nil
}

func (s *blogPostService) GetListOfBlogPosts(ctx context.Context, limit, offset int) ([]*dto.BlogPost, error) {
	var repoBlogPosts []*model.BlogPost

	var result *gorm.DB

	query := s.orm.WithContext(ctx).Offset(offset).Order("id ASC")

	result = query.Find(&repoBlogPosts)

	if result.Error != nil {
		return nil, result.Error
	}

	blogPosts := make([]*dto.BlogPost, len(repoBlogPosts))
	for i, blogPost := range repoBlogPosts {
		blogPosts[i] = &dto.BlogPost{
			ID:        blogPost.ID,
			CreatedAt: *blogPost.CreatedAt,
			UpdatedAt: *blogPost.UpdatedAt,
			BlogPostToCreate: dto.BlogPostToCreate{
				Title:   *blogPost.Title,
				Content: *blogPost.Content,
				Author:  *blogPost.Author,
			},
		}
	}

	return blogPosts, nil
}

func (s *blogPostService) UpdateBlogPost(ctx context.Context, postId int, updateBlogPost *dto.UpdateBlogPost) (*dto.BlogPost, error) {
	var updatedBlogPost model.BlogPost

	result := s.orm.WithContext(ctx).
		Model(&updatedBlogPost).
		Clauses(clause.Returning{}).
		Omit("ID", "created_at", "author").
		Where("id = ?", postId).
		Updates(&model.BlogPost{
			Title:   updateBlogPost.Title,
			Content: updateBlogPost.Content,
		})

	if result.RowsAffected == 0 {
		return nil, errors.New("blog post not found")
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog post not found")
		}

		return nil, result.Error
	}

	return &dto.BlogPost{
		ID:        updatedBlogPost.ID,
		CreatedAt: *updatedBlogPost.CreatedAt,
		UpdatedAt: *updatedBlogPost.UpdatedAt,
		BlogPostToCreate: dto.BlogPostToCreate{
			Title:   *updatedBlogPost.Title,
			Content: *updatedBlogPost.Content,
			Author:  *updatedBlogPost.Author,
		},
	}, nil
}

func (s *blogPostService) DeleteBlogPost(ctx context.Context, postId int) error {

	result := s.orm.WithContext(ctx).Delete(&model.BlogPost{}, postId)

	if result.Error != nil {
		if helper.ViolatesForeignKeyError(result.Error) {
			return errors.New("cannot be deleted because the database has data associated with it")
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("blog post not found")
	}

	return nil
}
