package middleware

import (
	"errors"
	"strings"

	"company-service/internal/helper"

	"github.com/labstack/echo/v5"
)

const (
	UserIDKey = "user_id"
	RoleKey   = "role"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.Unauthorized(c, "authorization header is required")
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.Unauthorized(c, "invalid authorization header")
		}

		tokenString := parts[1]

		claims, err := helper.ValidateToken(tokenString)
		if err != nil {
			return helper.Unauthorized(c, "invalid or expired token")
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(RoleKey, claims.Role)

		return next(c)
	}
}

func RequireRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			role, ok := c.Get(RoleKey).(string)

			if !ok {
				return helper.Unauthorized(c, "invalid user role")
			}

			if role != requiredRole {
				return helper.Forbidden(c, "access denied")
			}

			return next(c)
		}
	}
}

func GetUserID(c *echo.Context) (uint, error) {
	userID, ok := c.Get(UserIDKey).(uint)

	if !ok {
		return 0, errors.New("invalid user identity")
	}

	return userID, nil
}
