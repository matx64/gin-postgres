package db

import (
	"log"
	"os"

	"github.com/matx64/gin-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database. \n", err)
	}

	DB = db

	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.User{})
}
