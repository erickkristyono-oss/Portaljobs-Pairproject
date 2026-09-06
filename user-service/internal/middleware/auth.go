package middleware

import (
	"net/http"
	"strings"

	"user-service/internal/helper"

	"github.com/labstack/echo/v5"
)

const (
	UserIDKey = "user_id"
	RoleKey   = "role"
)

// JWTMiddleware digunakan untuk memvalidasi JWT token
// dan menyimpan UserID serta Role ke context.
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "authorization header is required",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "invalid authorization header",
			})
		}

		tokenString := strings.TrimSpace(parts[1])

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "token is required",
			})
		}

		claims, err := helper.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "invalid or expired token",
			})
		}

		// Simpan data dari JWT ke Echo Context
		c.Set(UserIDKey, claims.UserID)
		c.Set(RoleKey, claims.Role)

		return next(c)
	}
}

// RequireRole digunakan untuk membatasi endpoint
// berdasarkan role yang terdapat di JWT token.
func RequireRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// Ambil role dari context
			roleValue := c.Get(RoleKey)

			if roleValue == nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"responseCode":    http.StatusUnauthorized,
					"responseMessage": "role not found",
				})
			}

			// Pastikan role berupa string
			role, ok := roleValue.(string)

			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"responseCode":    http.StatusUnauthorized,
					"responseMessage": "invalid role",
				})
			}

			// Cek apakah role user sesuai dengan role endpoint
			if role != requiredRole {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"responseCode":    http.StatusForbidden,
					"responseMessage": "forbidden",
				})
			}

			// Role sesuai → lanjut ke handler
			return next(c)
		}
	}
}
