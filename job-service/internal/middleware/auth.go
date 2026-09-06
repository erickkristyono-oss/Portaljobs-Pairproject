package middleware

import (
	"net/http"
	"strings"

	"job-service/internal/helper"

	"github.com/labstack/echo/v5"
)

const (
	UserIDKey = "user_id"
	RoleKey   = "role"
)

// AuthMiddleware memvalidasi JWT dari Authorization header.
func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, helper.Response{
					ResponseCode:    http.StatusUnauthorized,
					ResponseMessage: "authorization header is required",
				})
			}

			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, helper.Response{
					ResponseCode:    http.StatusUnauthorized,
					ResponseMessage: "invalid authorization header",
				})
			}

			token := parts[1]

			claims, err := helper.ParseToken(token, jwtSecret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.Response{
					ResponseCode:    http.StatusUnauthorized,
					ResponseMessage: "invalid or expired token",
				})
			}

			if claims.UserID == 0 {
				return c.JSON(http.StatusUnauthorized, helper.Response{
					ResponseCode:    http.StatusUnauthorized,
					ResponseMessage: "invalid user id",
				})
			}

			// Simpan data JWT ke Echo context.
			c.Set(UserIDKey, claims.UserID)
			c.Set(RoleKey, claims.Role)

			return next(c)
		}
	}
}

// RequireRole membatasi endpoint berdasarkan role.
func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			currentRole, ok := c.Get(RoleKey).(string)

			if !ok || currentRole == "" {
				return c.JSON(http.StatusUnauthorized, helper.Response{
					ResponseCode:    http.StatusUnauthorized,
					ResponseMessage: "role not found",
				})
			}

			if currentRole != role {
				return c.JSON(http.StatusForbidden, helper.Response{
					ResponseCode:    http.StatusForbidden,
					ResponseMessage: "forbidden",
				})
			}

			return next(c)
		}
	}
}
