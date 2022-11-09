package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Title     string    `gorm:"type:text;not null" json:"title,omitempty"`
	Content   string    `gorm:"type:text;not null" json:"content,omitempty"`
	Image     string    `gorm:"type:text;not null" json:"image,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
