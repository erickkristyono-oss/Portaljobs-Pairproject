package router

import (
	"github.com/labstack/echo/v5"

	"company-service/internal/handler"
	"company-service/internal/middleware"
)

func RegisterRoutes(
	e *echo.Echo,
	companyHandler *handler.CompanyHandler,
	companyProfileHandler *handler.CompanyProfileHandler,
) {
	auth := e.Group("")

	auth.Use(middleware.JWTMiddleware)
	auth.Use(middleware.RequireRole("company"))

	// Company Profile
	auth.POST("/companies/profile", companyProfileHandler.Create)
	auth.GET("/companies/me/profile", companyProfileHandler.Get)
	auth.PUT("/companies/me/profile", companyProfileHandler.Update)

	// internal point
	e.GET("/companies/:id", companyHandler.GetByID)
	e.GET("/internal/companies/user/:user_id", companyProfileHandler.GetByUserID)
}
