package routes

import (
	"github.com/gin-gonic/gin"

	"gorm-research/controllers"
)

const (
	relativePath = "posts"
)

type PostRouteController struct {
	postController *controllers.PostController
}

func NewPostRouteController(postController *controllers.PostController) *PostRouteController {
	return &PostRouteController{postController: postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group(relativePath)
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.GetPosts)
	router.DELETE("/:id", pc.postController.DeletePost)
	router.PUT("/:id", pc.postController.UpdatePost)
}
