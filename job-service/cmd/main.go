package main

import (
	"log"

	"job-service/config"
	"job-service/internal/client"
	"job-service/internal/handler"
	"job-service/internal/helper"
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

	log.Println("database connected successfully")

	// Repository
	jobRepository := postgres.NewJobRepository(db)
	jobRequiredSkillRepository := postgres.NewJobRequiredSkillRepository(db)

	companyClient := client.NewCompanyClient(
		"http://localhost:8082",
	)
	// Usecase
	jobUsecase := job.NewJobUsecase(
		jobRepository,
		jobRequiredSkillRepository,
		companyClient,
	)

	// Handler
	jobHandler := handler.NewJobHandler(jobUsecase)

	// Echo
	e := echo.New()

	// Validator
	e.Validator = helper.NewValidator()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Router
	router.SetupRouter(
		e,
		jobHandler,
	)

	// Start Server
	const appPort = ":8083"

	log.Println("job-service running on http://localhost" + appPort)

	if err := e.Start(appPort); err != nil {
		log.Fatal("failed to start job-service:", err)
	}
}
