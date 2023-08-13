package model

import (
	"time"
)

type PostComment struct {
	ID        int        `gorm:"primaryKey,autoIncrement"`
	PostID    *int       `gorm:"column:post_id;type:integer;not null"`
	Author    *string    `gorm:"column:author;type:text;not null"`
	Content   *string    `gorm:"column:content;type:text;not null"`
	IsDeleted *bool      `gorm:"column:is_deleted;default:false"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;"`

	Post *BlogPost `gorm:"foreignKey:PostID;references:ID"`
}

func (PostComment) TableName() string {
	return "post_comment"
}
