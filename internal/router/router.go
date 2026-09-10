package router

import (
	"net/http"

	"portaljob/docs"
	"portaljob/internal/config"
	"portaljob/internal/handler"
	"portaljob/internal/middleware"
	"portaljob/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth        *handler.AuthHandler
	Job         *handler.JobHandler
	Company     *handler.CompanyHandler
	Profile     *handler.ProfileHandler
	Skill       *handler.SkillHandler
	Portfolio   *handler.PortfolioHandler
	Application *handler.ApplicationHandler
	Report      *handler.ReportHandler
	Admin       *handler.AdminHandler
	SkillTag    *handler.SkillTagHandler
}

func New(h Handlers, jwtManager *jwt.Manager, cfg *config.Config) *gin.Engine {
	r := gin.Default() // includes Logger + Recovery

	// global security middleware
	r.Use(middleware.SecureHeaders())
	r.Use(middleware.CORS(cfg.CORSAllowOrigin))
	r.Use(middleware.BodyLimit(1 << 20)) // 1 MB
	r.Use(middleware.RateLimit(cfg.RateLimitPerMin))

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	docs.Register(r)

	api := r.Group("/api/v1")

	// public
	api.POST("/register", h.Auth.Register)
	api.POST("/login", h.Auth.Login)
	api.GET("/jobs", h.Job.List)
	api.GET("/jobs/:id", h.Job.Detail)
	api.GET("/skill-tags", h.SkillTag.ListVocab)

	// authenticated (any role)
	authed := api.Group("")
	authed.Use(middleware.Auth(jwtManager))
	authed.POST("/reports", h.Report.Create)

	// jobseeker only
	js := authed.Group("")
	js.Use(middleware.RequireRole("jobseeker"))
	{
		js.GET("/me/profile", h.Profile.Get)
		js.PUT("/me/profile", h.Profile.Save)
		js.GET("/me/skills", h.Skill.List)
		js.POST("/me/skills", h.Skill.Add)
		js.DELETE("/me/skills/:id", h.Skill.Delete)
		js.GET("/me/portfolios", h.Portfolio.List)
		js.POST("/me/portfolios", h.Portfolio.Add)
		js.DELETE("/me/portfolios/:id", h.Portfolio.Delete)
		js.POST("/jobs/:id/apply", h.Application.Apply)
		js.GET("/me/applications", h.Application.ListMine)
		js.GET("/me/skill-tags", h.SkillTag.GetMine)
		js.PUT("/me/skill-tags", h.SkillTag.SetMine)
		js.GET("/jobs/:id/match", h.Application.JobMatch)
	}

	// company only
	co := authed.Group("")
	co.Use(middleware.RequireRole("company"))
	{
		co.GET("/company/profile", h.Company.Get)
		co.PUT("/company/profile", h.Company.Save)
		co.POST("/jobs", h.Job.Create)
		co.GET("/company/jobs", h.Job.ListMine)
		co.GET("/jobs/:id/applications", h.Application.ListApplicants)
		co.PATCH("/applications/:id", h.Application.Decide)
	}

	// admin only
	admin := authed.Group("/admin")
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.GET("/users", h.Admin.ListUsers)
		admin.PATCH("/users/:id/suspend", h.Admin.SuspendUser)
		admin.GET("/reports", h.Admin.ListReports)
		admin.PATCH("/reports/:id", h.Admin.ResolveReport)
	}

	return r
}
