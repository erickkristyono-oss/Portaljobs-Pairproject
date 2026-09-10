package router

import (
	"admin-service/internal/handler"
	"admin-service/internal/middleware"

	"github.com/labstack/echo/v5"
)

func SetupRouter(
	e *echo.Echo,
	reportHandler *handler.ReportHandler,
	userHandler *handler.UserHandler,
) {

	admin := e.Group(
		"/admin",
		middleware.JWTMiddleware,
		middleware.RequireRole("admin"),
	)

	// =========================
	// REPORT
	// =========================

	admin.POST("/reports", reportHandler.Create)
	admin.GET("/reports", reportHandler.FindAll)
	admin.GET("/reports/:id", reportHandler.FindByID)
	admin.PATCH("/reports/:id/status", reportHandler.UpdateStatus)

	admin.PATCH("/users/:id/status", userHandler.UpdateStatus)
}
