package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"application-service/config"
	"application-service/internal/client"
	"application-service/internal/handler"
	"application-service/internal/helper"
	"application-service/internal/notification"
	"application-service/internal/repository/postgres"
	"application-service/internal/router"
	"application-service/internal/skilltag"
	application "application-service/internal/usecase/app"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// DATABASE
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("database connected successfully")

	// REPOSITORY
	applicationRepository := postgres.NewApplicationRepository(db)

	// SERVICE CLIENTS

	// User Service :8081
	userClient := client.NewUserClient(
		"http://localhost:8081",
	)

	// Company Service :8082
	companyClient := client.NewCompanyClient(
		"http://localhost:8082",
	)

	// Job Service :8083
	jobClient := client.NewJobClient(
		"http://localhost:8083",
	)

	// SKILL TAG USECASE
	skillTagUsecase := skilltag.NewSkillTagUsecase()

	// FONNTE WHATSAPP API
	fonnteAPIURL := os.Getenv("FONNTE_API_URL")
	fonnteToken := os.Getenv("FONNTE_TOKEN")

	fonnteClient := client.NewFonnteClient(
		fonnteAPIURL,
		fonnteToken,
	)

	notificationService := notification.NewWhatsAppNotification(
		fonnteClient,
	)

	fmt.Println("FONNTE_API_URL:", fonnteAPIURL)
	fmt.Println("FONNTE_TOKEN SET:", fonnteToken != "")

	// USECASE
	applicationUsecase := application.NewApplicationUsecase(
		applicationRepository,
		jobClient,
		userClient,
		companyClient,
		skillTagUsecase,
		notificationService,
	)

	// HANDLER
	applicationHandler := handler.NewApplicationHandler(
		applicationUsecase,
	)

	// ECHO
	e := echo.New()

	// Validator
	e.Validator = helper.NewValidator()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// ROUTER
	router.SetupRouter(
		e,
		applicationHandler,
	)

	// SERVER
	const appPort = ":8084"

	log.Println(
		"application-service running on http://localhost" + appPort,
	)

	if err := e.Start(appPort); err != nil {
		log.Fatal("failed to start application-service:", err)
	}
}
