package model

import "time"

type BlogPost struct {
	ID        int        `gorm:"primaryKey,autoIncrement"`
	Title     *string    `gorm:"column:title;type:text;not null"`
	Content   *string    `gorm:"column:content;type:text;not null"`
	Author    *string    `gorm:"column:author;type:text;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;"`
}

func (u *BlogPost) TableName() string {
	return "blog_post"
}
