package router

import (
	"job-service/internal/handler"
	"job-service/internal/middleware"

	"github.com/labstack/echo/v5"
)

func SetupRouter(
	e *echo.Echo,
	jobHandler *handler.JobHandler,
) {
	// Semua user yang sudah login
	// dapat melihat job yang published.
	e.GET("/jobs", jobHandler.GetAll, middleware.JWTMiddleware)
	e.GET("/jobs/:id", jobHandler.GetByID, middleware.JWTMiddleware)

	//khusus company.
	e.POST("/jobs", jobHandler.Create, middleware.JWTMiddleware, middleware.RequireRole("company"))
	e.PUT("/jobs/:id", jobHandler.Update, middleware.JWTMiddleware, middleware.RequireRole("company"))
	e.DELETE("/jobs/:id", jobHandler.Delete, middleware.JWTMiddleware, middleware.RequireRole("company"))
	e.PATCH("/jobs/:id/publish", jobHandler.Publish, middleware.JWTMiddleware, middleware.RequireRole("company"))
	e.PATCH("/jobs/:id/close", jobHandler.Close, middleware.JWTMiddleware, middleware.RequireRole("company"))

	// untuk mengambil required skill dari job
	e.GET("/internal/jobs/:id/required-skills", jobHandler.GetRequiredSkills)
}
