package handler

import (
	"net/http"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct{ reports *usecase.ReportUsecase }

func NewReportHandler(reports *usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{reports: reports}
}

// Create — POST /api/v1/reports (any authenticated user).
func (h *ReportHandler) Create(c *gin.Context) {
	var req ReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	report, err := h.reports.Create(c.Request.Context(), middleware.UserID(c), req.TargetType, req.TargetID, req.Reason)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "report filed", report)
}
