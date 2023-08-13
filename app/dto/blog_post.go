package dto

import "time"

type BlogPost struct {
	ID int `json:"id,omitempty"`
	BlogPostToCreate
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BlogPostToCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type UpdateBlogPost struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
