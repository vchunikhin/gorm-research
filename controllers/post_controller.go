package controllers

import (
	"net/http"
	"strconv"
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
			ctx.JSON(http.StatusConflict, failResponse("Post with that title already exists"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(result.Error.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, successResponse(post))
}

func (pc PostController) GetPosts(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failResponse("Failed to convert a page value"))
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failResponse("Failed to convert a limit value"))
		return
	}

	offset := (intPage - 1) * intLimit

	var posts []*models.Post
	result := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(result.Error.Error()))
		return
	}
	ctx.JSON(http.StatusOK, paginationResponse(len(posts), posts))
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	result := pc.DB.Delete(&models.Post{}, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(result.Error.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload *models.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, failResponse(err.Error()))
		return
	}

	var post *models.Post
	result := pc.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(result.Error.Error()))
		return
	}

	now := time.Now()

	updatedPost := &models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		CreatedAt: post.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&post).Updates(updatedPost)

	ctx.JSON(http.StatusOK, successResponse(post))
}
