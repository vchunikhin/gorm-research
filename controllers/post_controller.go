package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gorm-research/internal/models"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) *PostController {
	return &PostController{DB: DB}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	post := &models.Post{
		ID:        uuid.New(),
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(post)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": post})
}
