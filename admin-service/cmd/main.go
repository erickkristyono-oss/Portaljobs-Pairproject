package main

import (
	"log"

	"admin-service/config"
	"admin-service/internal/handler"
	"admin-service/internal/repository/postgres"
	"admin-service/internal/router"
	"admin-service/internal/usecase/report"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	// Load environment variables dari .env
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// DATABASE
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// REPOSITORY
	reportRepository := postgres.NewReportRepository(db)

	// USECASE
	reportUsecase := report.NewReportUsecase(reportRepository)

	// HANDLER
	reportHandler := handler.NewReportHandler(reportUsecase)

	// ECHO
	e := echo.New()

	// Request logger
	e.Use(middleware.RequestLogger())

	// Recover dari panic
	e.Use(middleware.Recover())

	// ROUTER
	router.SetupRouter(
		e,
		reportHandler,
	)

	// SERVER

	log.Println(
		"admin-service running on http://localhost:8084",
	)

	if err := e.Start(":8084"); err != nil {
		log.Fatal(err)
	}
}
