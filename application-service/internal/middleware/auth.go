package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"

	"application-service/internal/helper"
)

const (
	UserIDKey        = "user_id"
	RoleKey          = "role"
	AuthorizationKey = "authorization"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c *echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.Unauthorized(
				c,
				"authorization header required",
			)
		}

		parts := strings.SplitN(
			authHeader,
			" ",
			2,
		)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {

			return helper.Unauthorized(
				c,
				"invalid authorization header",
			)
		}

		tokenString := parts[1]

		claims := &helper.JWTClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {

				if token.Method != jwt.SigningMethodHS256 {
					return nil, errors.New(
						"unexpected signing method",
					)
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return helper.Unauthorized(
				c,
				"invalid or expired token",
			)
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(RoleKey, claims.Role)

		return next(c)
	}
}

func RequireRole(role string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			currentRole, ok := c.Get(RoleKey).(string)

			if !ok || currentRole != role {
				return helper.Forbidden(
					c,
					"forbidden",
				)
			}

			return next(c)
		}
	}
}
