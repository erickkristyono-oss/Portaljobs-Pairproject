package main

import (
	"log"
	"os"

	"admin-service/config"
	"admin-service/internal/client"
	"admin-service/internal/handler"
	"admin-service/internal/helper"
	"admin-service/internal/repository/postgres"
	"admin-service/internal/router"
	"admin-service/internal/usecase/report"
	userusecase "admin-service/internal/usecase/user"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	// Load environment variables dari .env
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// =========================
	// DATABASE
	// =========================

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// =========================
	// REPOSITORY
	// =========================

	reportRepository := postgres.NewReportRepository(db)

	// =========================
	// REPORT USECASE
	// =========================

	reportUsecase := report.NewReportUsecase(
		reportRepository,
	)

	// =========================
	// USER CLIENT
	// =========================

	userClient := client.NewUserClient(
		os.Getenv("USER_SERVICE_URL"),
	)

	// =========================
	// USER USECASE
	// =========================

	userUsecase := userusecase.NewUserUsecase(
		userClient,
	)

	// =========================
	// HANDLER
	// =========================

	reportHandler := handler.NewReportHandler(
		reportUsecase,
	)

	userHandler := handler.NewUserHandler(
		userUsecase,
	)

	// =========================
	// ECHO
	// =========================

	e := echo.New()

	// Validator
	e.Validator = helper.NewValidator()

	// Request logger
	e.Use(middleware.RequestLogger())

	// Recover dari panic
	e.Use(middleware.Recover())

	// =========================
	// ROUTER
	// =========================

	router.SetupRouter(
		e,
		reportHandler,
		userHandler,
	)

	// =========================
	// SERVER
	// =========================

	const appPort = ":8085"

	log.Println(
		"admin-service running on http://localhost:8085",
	)

	if err := e.Start(appPort); err != nil {
		log.Fatal(err)
	}
}
