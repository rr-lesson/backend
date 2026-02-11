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
		AllowOrigins: "*",
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

	classRepo := repositories.NewClassRepository(db)
	subjectRepo := repositories.NewSubjectRepository(db)
	lessonRepo := repositories.NewLessonRepository(db)
	videoRepo := repositories.NewVideoRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)

	classHandler := handlers.NewClassHandler(classRepo)
	subjectHandler := handlers.NewSubjectHandler(subjectRepo)
	lessonHandler := handlers.NewLessonHandler(lessonRepo)
	videoHandler := handlers.NewVideoHandler(minio, videoRepo)
	questionHandler := handlers.NewQuestionHandler(questionRepo)

	classHandler.RegisterRoutes(v1)
	subjectHandler.RegisterRoutes(v1)
	lessonHandler.RegisterRoutes(v1)
	videoHandler.RegisterRoutes(v1)
	questionHandler.RegisterRoutes(v1)

	log.Printf("Server running on port %s", "8080")
	log.Fatal(app.Listen(":" + "8080"))
}
