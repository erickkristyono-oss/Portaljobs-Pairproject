package router

import (
	"job-service/internal/handler"
	"job-service/internal/middleware"

	"github.com/labstack/echo/v5"
)

func SetupRouter(
	e *echo.Echo,
	jobHandler *handler.JobHandler,
	jwtSecret string,
) {
	// Public routes

	// Semua user dapat melihat daftar job.
	e.GET("/jobs", jobHandler.GetAll)
	e.GET("/jobs/:id", jobHandler.GetByID)

	// Company routes

	company := e.Group(
		"",
		middleware.AuthMiddleware(jwtSecret),
		middleware.RequireRole("company"),
	)

	company.POST("/jobs", jobHandler.Create)
	company.PUT("/jobs/:id", jobHandler.Update)
	company.DELETE("/jobs/:id", jobHandler.Delete)
	company.PATCH("/jobs/:id/publish", jobHandler.Publish)
	company.PATCH("/jobs/:id/close", jobHandler.Close)
}
