package main

import (
	"log"

	"job-service/config"
	"job-service/internal/handler"
	"job-service/internal/repository/postgres"
	"job-service/internal/router"
	"job-service/internal/usecase/job"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// Connect Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Repository

	// Repository untuk jobs
	jobRepository := postgres.NewJobRepository(db)

	// Repository untuk required skills
	jobRequiredSkillRepository := postgres.NewJobRequiredSkillRepository(db)

	// Usecase
	jobUsecase := job.NewJobUsecase(jobRepository, jobRequiredSkillRepository)

	// Handler
	jobHandler := handler.NewJobHandler(jobUsecase)

	// Echo
	e := echo.New()

	// Global Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Router
	router.SetupRouter(
		e,
		jobHandler,
	)

	// Start Server
	log.Println("job-service running on http://localhost:8082")

	if err := e.Start(":8082"); err != nil {
		log.Fatal("failed to start job-service:", err)
	}
}
