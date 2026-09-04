package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct{ apps *usecase.ApplicationUsecase }

func NewApplicationHandler(apps *usecase.ApplicationUsecase) *ApplicationHandler {
	return &ApplicationHandler{apps: apps}
}

// Apply — POST /api/v1/jobs/:id/apply (jobseeker).
func (h *ApplicationHandler) Apply(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid job id")
		return
	}
	app, err := h.apps.Apply(c.Request.Context(), middleware.UserID(c), uint(jobID))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "application submitted", app)
}

// ListMine — GET /api/v1/me/applications (jobseeker).
func (h *ApplicationHandler) ListMine(c *gin.Context) {
	apps, err := h.apps.ListMine(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "my applications", apps)
}

// ListApplicants — GET /api/v1/jobs/:id/applications (company owner).
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

// Decide — PATCH /api/v1/applications/:id (company owner).
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
