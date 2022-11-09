package main

import (
	"log"

	"gorm-research/config"
	"gorm-research/internal/db"
	"gorm-research/internal/models"
)

const (
	configPath = "."
)

var (
	configVal *config.Config
)

func init() {
	var err error
	configVal, err = config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Couldn't load environment variables", err)
	}
	db.ConnectDB(configVal)
}

func main() {
	if err := db.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatal("Failed to migrate data", err.Error())
	}
	log.Println("Migration complete")
}
