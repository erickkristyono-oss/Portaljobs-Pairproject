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

	applications := e.Group(
		"/applications",
		middleware.JWTMiddleware,
	)

	// Jobseeker
	applications.POST("", applicationHandler.Apply, middleware.RequireRole("jobseeker"))
	applications.GET("/me", applicationHandler.GetMyApplications, middleware.RequireRole("jobseeker"))
	applications.GET("/:id", applicationHandler.GetByID, middleware.RequireRole("jobseeker"))
	applications.PATCH("/:id/status", applicationHandler.UpdateStatus, middleware.RequireRole("company"))
	applications.DELETE("/:id", applicationHandler.Delete, middleware.RequireRole("jobseeker"))

	// Company melihat semua applicant
	e.GET("/jobs/:job_id/applications",
		applicationHandler.GetApplicationsByJobID,
		middleware.JWTMiddleware,
		middleware.RequireRole("company"),
	)
}
