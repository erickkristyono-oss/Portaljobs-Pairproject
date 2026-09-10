package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v5"
)

func createProxy(target string) echo.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return echo.WrapHandler(reverseProxy)
}

func RegisterProxyRoutes(e *echo.Echo) {

	// ============================================================
	// USER SERVICE
	// ============================================================

	userProxy := createProxy("http://localhost:8081")

	e.Any("/register", userProxy)
	e.Any("/login", userProxy)

	e.Any("/me/*", userProxy)

	e.Any("/internal/users/*", userProxy)

	// ============================================================
	// COMPANY SERVICE
	// ============================================================

	companyProxy := createProxy("http://localhost:8082")

	e.Any("/companies", companyProxy)
	e.Any("/companies/*", companyProxy)

	// ============================================================
	// JOB SERVICE
	// ============================================================

	jobProxy := createProxy("http://localhost:8083")

	e.Any("/jobs", jobProxy)
	e.Any("/jobs/*", jobProxy)

	// ============================================================
	// APPLICATION SERVICE
	// ============================================================

	applicationProxy := createProxy("http://localhost:8084")

	e.Any("/applications", applicationProxy)
	e.Any("/applications/*", applicationProxy)

	// ============================================================
	// ADMIN SERVICE
	// ============================================================

	adminProxy := createProxy("http://localhost:8085")

	e.Any("/admin", adminProxy)
	e.Any("/admin/*", adminProxy)
}
