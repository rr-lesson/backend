package main

import (
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/pkg/database"
	"backend/pkg/minio"
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// @title 			RR-Lesson API
// @version 		1.0
// @description Backend API for RR Lesson Application
// @host 				localhost:8080
// @BasePath 		/
func main() {
	if os.Getenv("GO_ENV") != "production" {
		_ = godotenv.Load()
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: func() string {
			if os.Getenv("GO_ENV") == "production" {
				return "https://api-rr.rizalanggoro.my.id"
			} else {
				return "http://rizalanggoro:3000"
			}
		}(),
		AllowCredentials: true,
	}))
	if os.Getenv("GO_ENV") != "production" {
		app.Use(swagger.New(swagger.Config{
			FilePath: "./docs/swagger.json",
			BasePath: "/api/v1",
			Path:     "docs",
		}))
	}

	v1 := app.Group("/api/v1")

	db := database.New()
	minio := minio.New()

	authRepo := repositories.NewAuthRepository(db)
	classRepo := repositories.NewClassRepository(db)
	subjectRepo := repositories.NewSubjectRepository(db)
	lessonRepo := repositories.NewLessonRepository(db)
	videoRepo := repositories.NewVideoRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)

	authHandler := handlers.NewAuthHandler(authRepo)
	classHandler := handlers.NewClassHandler(classRepo)
	subjectHandler := handlers.NewSubjectHandler(subjectRepo)
	lessonHandler := handlers.NewLessonHandler(lessonRepo)
	videoHandler := handlers.NewVideoHandler(minio, videoRepo)
	questionHandler := handlers.NewQuestionHandler(questionRepo)

	authHandler.RegisterRoutes(v1)
	classHandler.RegisterRoutes(v1)
	subjectHandler.RegisterRoutes(v1)
	lessonHandler.RegisterRoutes(v1)
	videoHandler.RegisterRoutes(v1)
	questionHandler.RegisterRoutes(v1)

	log.Printf("Server running on port %s", "8080")
	log.Fatal(app.Listen(":" + "8080"))
}
