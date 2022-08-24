package db

import (
	"log"

	"github.com/sumirart/go-gin-gorm-books/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error opening database", err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
