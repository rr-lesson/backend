package database

import (
	"backend/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" port=" + dbPort +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.UserSession{},
		&models.Class{},
		&models.Subject{},
		&models.Lesson{},
		&models.Video{},
		&models.Question{},
		&models.QuestionAttachment{},
		&models.Answer{},
	); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
