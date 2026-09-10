package router

import (
	"github.com/labstack/echo/v5"

	"user-service/internal/handler"
	"user-service/internal/middleware"
)

func RegisterRoutes(
	e *echo.Echo,
	userHandler *handler.UserHandler,
	profileHandler *handler.ProfileHandler,
	skillHandler *handler.SkillHandler,
	portfolioHandler *handler.PortfolioHandler,
) {
	// Public Routes

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	// Internal routes
	e.PATCH("/internal/users/:id/status", userHandler.UpdateStatus)

	// Authenticated Routes

	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware)

	// Jobseeker Routes

	jobseeker := auth.Group("")
	jobseeker.Use(middleware.RequireRole("jobseeker"))

	// Profile
	jobseeker.POST("/me/profile", profileHandler.Create)
	jobseeker.GET("/me/profile", profileHandler.GetByUserID)
	e.GET("/internal/users/:user_id/profile", profileHandler.GetByUserIDInternal)
	e.GET("/internal/users/:user_id/skills", skillHandler.GetByUserIDInternal)
	jobseeker.PUT("/me/profile", profileHandler.Update)

	// Skill
	jobseeker.GET("/me/skills", skillHandler.GetByUserID)
	jobseeker.POST("/me/skills", skillHandler.Create)
	jobseeker.PUT("/me/skills/:id", skillHandler.Update)
	jobseeker.DELETE("/me/skills/:id", skillHandler.Delete)

	// Portfolio
	jobseeker.GET("/me/portfolios", portfolioHandler.GetByUserID)
	jobseeker.GET("/me/portfolios/:id", portfolioHandler.GetByID)
	jobseeker.POST("/me/portfolios", portfolioHandler.Create)
	jobseeker.PUT("/me/portfolios/:id", portfolioHandler.Update)
	jobseeker.DELETE("/me/portfolios/:id", portfolioHandler.Delete)
}
