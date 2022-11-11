package models

import (
	"time"

	"github.com/google/uuid"

	"gorm-research/internal/types"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Title     string    `gorm:"type:text;not null"`
	Content   string    `gorm:"type:text;not null"`
	Image     string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (p *Post) ToResponse() *PostResponse {
	return &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		Image:     p.Image,
		CreatedAt: types.CustomTime{Time: p.CreatedAt},
		UpdatedAt: types.CustomTime{Time: p.UpdatedAt},
	}
}

type PostResponse struct {
	ID        uuid.UUID        `json:"id,omitempty"`
	Title     string           `json:"title,omitempty"`
	Content   string           `json:"content,omitempty"`
	Image     string           `json:"image,omitempty"`
	CreatedAt types.CustomTime `json:"created_at,omitempty"`
	UpdatedAt types.CustomTime `json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePostRequest struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
