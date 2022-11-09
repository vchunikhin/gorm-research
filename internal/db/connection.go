package db

//mysql driver: https://github.com/go-sql-driver/mysql#parameters

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm-research/config"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBUserPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	log.Println("Connected Successfully to the Database")
}
