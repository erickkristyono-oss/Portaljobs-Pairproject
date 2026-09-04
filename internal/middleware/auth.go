package middleware

import (
	"net/http"
	"strings"

	"portaljob/pkg/jwt"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserID = "user_id"
	ctxRole   = "role"
)

// Auth verifies the Bearer JWT and stores user_id + role in the gin context.
func Auth(jwtManager *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "missing or malformed token")
			c.Abort()
			return
		}
		claims, err := jwtManager.Parse(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}
		c.Set(ctxUserID, claims.UserID)
		c.Set(ctxRole, claims.Role)
		c.Next()
	}
}

// RequireRole allows the request only if the caller has one of the given roles.
// Use it after Auth, e.g. RequireRole("company").
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(ctxRole)
		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}
		response.Error(c, http.StatusForbidden, "you do not have access to this resource")
		c.Abort()
	}
}

// UserID reads the authenticated user's id from the context.
func UserID(c *gin.Context) uint {
	if v, ok := c.Get(ctxUserID); ok {
		if id, ok := v.(uint); ok {
			return id
		}
	}
	return 0
}

// CurrentRole reads the authenticated user's role from the context.
func CurrentRole(c *gin.Context) string {
	return c.GetString(ctxRole)
}
