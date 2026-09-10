package main

import (
	"github.com/labstack/echo/v5"

	"gateway/proxy"
)

func main() {
	e := echo.New()

	// Register reverse proxy
	proxy.RegisterProxyRoutes(e)

	// Centralized Swagger
	e.Static("/docs", "../docs")

	e.Logger.Info("Gateway running on :8079")

	e.Start(":8079")
}
