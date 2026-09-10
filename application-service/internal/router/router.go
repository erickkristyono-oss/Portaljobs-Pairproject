package router

import (
	"github.com/labstack/echo/v5"

	"application-service/internal/handler"
	"application-service/internal/middleware"
)

func SetupRouter(
	e *echo.Echo,
	applicationHandler *handler.ApplicationHandler,
) {
	// Application

	applications := e.Group(
		"/applications",
		middleware.JWTMiddleware,
	)

	// Jobseeker

	jobseeker := applications.Group("", middleware.RequireRole("jobseeker"))

	// Apply job
	jobseeker.POST("", applicationHandler.Apply)

	// Melihat application milik sendiri
	jobseeker.GET("/me", applicationHandler.GetMyApplications)

	// Melihat detail application
	jobseeker.GET("/:id", applicationHandler.GetByID)

	// Menghapus application
	jobseeker.DELETE("/:id", applicationHandler.Delete)

	// Company

	company := applications.Group(
		"/company",
		middleware.RequireRole("company"),
	)

	// Update status application
	company.PATCH("/:id/status", applicationHandler.UpdateStatus)

	// Company - melihat applicant berdasarkan job

	e.GET("/jobs/:job_id/applications",
		applicationHandler.GetApplicationsByJobID,
		middleware.JWTMiddleware,
		middleware.RequireRole("company"),
	)
}
