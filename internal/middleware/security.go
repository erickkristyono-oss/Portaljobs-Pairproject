package middleware

import (
	"net/http"
	"sync"
	"time"

	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

// CORS sets permissive-but-configurable CORS headers and short-circuits the
// preflight OPTIONS request.
func CORS(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// SecureHeaders adds a few standard hardening response headers.
func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
	}
}

// BodyLimit rejects request bodies larger than max bytes (basic DoS guard).
func BodyLimit(max int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, max)
		c.Next()
	}
}

// --- simple in-memory per-IP rate limiter (fixed 1-minute window) ---
// NOTE: the visitors map grows over time; for production add periodic cleanup
// or swap for a library / Redis. See README "kendala" notes.

type rlVisitor struct {
	count       int
	windowStart time.Time
}

var (
	rlMu       sync.Mutex
	rlVisitors = map[string]*rlVisitor{}
)

func RateLimit(perMin int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if perMin <= 0 {
			c.Next()
			return
		}
		ip := c.ClientIP()
		now := time.Now()

		rlMu.Lock()
		v, ok := rlVisitors[ip]
		if !ok || now.Sub(v.windowStart) > time.Minute {
			rlVisitors[ip] = &rlVisitor{count: 1, windowStart: now}
			rlMu.Unlock()
			c.Next()
			return
		}
		v.count++
		over := v.count > perMin
		rlMu.Unlock()

		if over {
			response.Error(c, http.StatusTooManyRequests, "rate limit exceeded, try again later")
			c.Abort()
			return
		}
		c.Next()
	}
}
