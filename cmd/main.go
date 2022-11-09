package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm-research/config"
	"gorm-research/controllers"
	"gorm-research/internal/db"
	"gorm-research/routes"
)

const (
	configPath   = "."
	relativePath = "/api"
	healthPath   = "/health"
)

var (
	server    *gin.Engine
	configVal *config.Config

	postController      *controllers.PostController
	postRouteController *routes.PostRouteController
)

func init() {
	var err error
	configVal, err = config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Couldn't load environment variables", err)
	}
	db.ConnectDB(configVal)

	postController = controllers.NewPostController(db.DB)
	postRouteController = routes.NewPostRouteController(postController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	server.Use(
		cors.New(corsConfig),
	)
	router := server.Group(relativePath)
	router.GET(healthPath, func(ctx *gin.Context) {
		msg := "OK"
		ctx.JSON(http.StatusOK, msg)
	})

	postRouteController.PostRoute(router)

	addr := fmt.Sprintf(":%s", configVal.ServerPort)
	if err := server.Run(addr); err != nil {
		log.Fatal(err)
	}
}
