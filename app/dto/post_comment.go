package dto

import "time"

type PostComment struct {
	ID int `json:"id,omitempty"`
	PostCommentToCreate
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PostCommentToCreate struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type DeletedCommentsStats struct {
	DeletedCount int `json:"deleted_count"`
	TotalCount   int `json:"total_count"`
}
