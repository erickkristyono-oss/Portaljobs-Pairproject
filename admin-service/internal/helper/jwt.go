package helper

import (
	"net/http"

	"admin-service/internal/middleware"

	"github.com/labstack/echo/v5"
)

func GetUserID(c *echo.Context) (uint, error) {

	userID, ok := c.Get(
		middleware.UserIDKey,
	).(uint)

	if !ok {
		return 0, echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	return userID, nil
}

func GetUserRole(c *echo.Context) (string, error) {

	role, ok := c.Get(
		middleware.RoleKey,
	).(string)

	if !ok {
		return "", echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	return role, nil
}
