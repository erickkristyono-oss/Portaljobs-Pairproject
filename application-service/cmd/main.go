package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"

	"application-service/config"
	"application-service/internal/client"
	"application-service/internal/handler"
	"application-service/internal/helper"
	"application-service/internal/repository/postgres"
	"application-service/internal/router"
	application "application-service/internal/usecase/app"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// Connect database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("database connected successfully")

	// Repository
	applicationRepository := postgres.NewApplicationRepository(db)

	// Service clients
	jobClient := client.NewJobClient(
		"http://localhost:8082",
	)

	companyClient := client.NewCompanyClient(
		"http://localhost:8081",
	)

	// Usecase
	applicationUsecase := application.NewApplicationUsecase(applicationRepository, jobClient, companyClient)

	// Handler
	applicationHandler := handler.NewApplicationHandler(applicationUsecase)

	// Echo
	e := echo.New()

	// Validator
	e.Validator = helper.NewValidator()

	// Router
	router.SetupRouter(
		e,
		applicationHandler,
	)

	// Application Service Port
	port := "8083"

	log.Println("application-service running on port", port)

	// Start server
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
