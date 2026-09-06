package router

import (
	"company-service/internal/handler"
	"company-service/internal/middleware"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(
	e *echo.Echo,
	companyHandler *handler.CompanyHandler,
) {
	// Public Routes

	// Melihat detail company berdasarkan ID.
	e.GET("/companies/:id", companyHandler.GetByID)

	// Protected Routes
	auth := e.Group("")

	// JWT authentication
	auth.Use(middleware.JWTMiddleware)

	// Company Routes

	company := auth.Group("")

	// Hanya user dengan role "company"
	// yang dapat mengakses endpoint.
	company.Use(middleware.RequireRole("company"))

	// Membuat company profile.
	company.POST("/me/company", companyHandler.Create)
	company.GET("/me/company", companyHandler.GetByUserID)
	company.PUT("/me/company", companyHandler.Update)

	internalRouter := e.Group("/internal")

	internalRouter.GET("/companies/user/:user_id", companyHandler.GetCompanyByUserID)

}
