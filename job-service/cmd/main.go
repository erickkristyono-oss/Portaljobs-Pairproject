package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"

	"job-service/config"
	"job-service/internal/client"
	"job-service/internal/handler"
	"job-service/internal/helper"
	"job-service/internal/repository/postgres"
	"job-service/internal/router"
	jobUsecase "job-service/internal/usecase/job"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// Connect PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Repository
	jobRepository := postgres.NewJobRepository(db)

	// Company service client
	companyClient := client.NewCompanyClient(
		"http://localhost:8081",
	)

	// Usecase
	jobUsecase := jobUsecase.NewJobUsecase(
		jobRepository,
		companyClient,
	)

	// Handler
	jobHandler := handler.NewJobHandler(jobUsecase)

	// Echo
	e := echo.New()

	//calidator
	e.Validator = helper.NewValidator()
	// Router
	router.SetupRouter(
		e,
		jobHandler,
		os.Getenv("JWT_SECRET"),
	)

	// Start server
	log.Println("job-service running on :8082")

	if err := e.Start(":8082"); err != nil {
		log.Fatal(err)
	}
}
