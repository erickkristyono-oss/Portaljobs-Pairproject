package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return echo.NewHTTPError(
				401,
				"authorization header is required",
			)
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(
				401,
				"invalid authorization header",
			)
		}

		tokenString := parts[1]

		// JWT Secret
		secret := os.Getenv("JWT_SECRET")

		if secret == "" {
			return echo.NewHTTPError(
				500,
				"JWT_SECRET is not configured",
			)
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {

				if token.Method != jwt.SigningMethodHS256 {
					return nil, errors.New("invalid signing method")
				}

				return []byte(secret), nil
			},
		)

		if err != nil {
			return echo.NewHTTPError(
				401,
				"invalid token",
			)
		}

		claims, ok := token.Claims.(*Claims)

		if !ok || !token.Valid {
			return echo.NewHTTPError(
				401,
				"invalid token",
			)
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		return next(c)
	}
}

func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			userRole, ok := c.Get("role").(string)

			if !ok || userRole != role {
				return echo.NewHTTPError(
					403,
					"forbidden",
				)
			}

			return next(c)
		}
	}
}
