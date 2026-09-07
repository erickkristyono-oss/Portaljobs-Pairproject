package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	apps  *usecase.ApplicationUsecase
	match *usecase.MatchUsecase
}

func NewApplicationHandler(apps *usecase.ApplicationUsecase, match *usecase.MatchUsecase) *ApplicationHandler {
	return &ApplicationHandler{apps: apps, match: match}
}

// Apply — POST /api/v1/jobs/:id/apply (jobseeker). Response includes a
// non-blocking skill-match summary (warning if fit is low).
func (h *ApplicationHandler) Apply(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid job id")
		return
	}
	userID := middleware.UserID(c)
	app, err := h.apps.Apply(c.Request.Context(), userID, uint(jobID))
	if err != nil {
		fail(c, err)
		return
	}
	// best-effort: attach match info; never fail the apply if match errors
	match, _ := h.match.Match(c.Request.Context(), userID, uint(jobID))
	response.OK(c, http.StatusCreated, "application submitted", gin.H{
		"application": app,
		"match":       match,
	})
}

// JobMatch — GET /api/v1/jobs/:id/match (jobseeker): fit before applying.
func (h *ApplicationHandler) JobMatch(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid job id")
		return
	}
	res, err := h.match.Match(c.Request.Context(), middleware.UserID(c), uint(jobID))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "match result", res)
}

func (h *ApplicationHandler) ListMine(c *gin.Context) {
	apps, err := h.apps.ListMine(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "my applications", apps)
}

func (h *ApplicationHandler) ListApplicants(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid job id")
		return
	}
	apps, err := h.apps.ListApplicants(c.Request.Context(), middleware.UserID(c), uint(jobID))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "applicants", apps)
}

func (h *ApplicationHandler) Decide(c *gin.Context) {
	appID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid application id")
		return
	}
	var req DecideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	app, err := h.apps.Decide(c.Request.Context(), middleware.UserID(c), uint(appID), req.Status)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "application updated", app)
}
