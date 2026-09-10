package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"

	"company-service/config"
	"company-service/internal/handler"
	"company-service/internal/repository/postgres"
	"company-service/internal/router"

	usecase_company "company-service/internal/usecase/company"
	usecase_profile "company-service/internal/usecase/company_profile"
)

func main() {

	// Load environment variable
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("database connected successfully")

	// Repository
	companyRepository := postgres.NewCompanyRepository(db)

	// Usecase
	profileUsecase := usecase_profile.NewCompanyProfileUsecase(
		companyRepository,
	)
	companyUsecase := usecase_company.NewCompanyUsecase(
		companyRepository,
	)

	// Handler
	profileHandler := handler.NewCompanyProfileHandler(
		profileUsecase,
	)
	companyHandler := handler.NewCompanyHandler(
		companyUsecase,
	)

	// Echo
	e := echo.New()

	// Router
	router.RegisterRoutes(
		e,
		companyHandler,
		profileHandler,
	)

	// Server
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8082"
	}

	log.Println("company-service running on http://localhost:" + port)

	if err := e.Start(":" + port); err != nil {
		log.Fatal("failed to start company-service:", err)
	}
}
