package main

import (
	"log"

	"company-service/config"
	"company-service/internal/handler"
	companyRepository "company-service/internal/repository/postgres"
	"company-service/internal/router"
	companyUsecase "company-service/internal/usecase/company"

	"github.com/labstack/echo/v5"
)

func main() {
	// Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Repository
	companyRepo := companyRepository.NewCompanyRepository(db)

	// Usecase
	companyUC := companyUsecase.NewCompanyUsecase(companyRepo)

	// Handler
	companyHandler := handler.NewCompanyHandler(companyUC)

	// Echo
	e := echo.New()

	// Router
	router.RegisterRoutes(
		e,
		companyHandler,
	)

	// Start Server
	const appPort = ":8081"

	log.Println("company-service running on", appPort)

	if err := e.Start(appPort); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
