package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct{ admin *usecase.AdminUsecase }

func NewAdminHandler(admin *usecase.AdminUsecase) *AdminHandler { return &AdminHandler{admin: admin} }

func (h *AdminHandler) ListUsers(c *gin.Context) {
	users, err := h.admin.ListUsers(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "users", users)
}

func (h *AdminHandler) SuspendUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}
	if err := h.admin.SuspendUser(c.Request.Context(), uint(id)); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "user suspended", nil)
}

func (h *AdminHandler) ListReports(c *gin.Context) {
	reports, err := h.admin.ListReports(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "reports", reports)
}

func (h *AdminHandler) ResolveReport(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid report id")
		return
	}
	var req ResolveReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	report, err := h.admin.ResolveReport(c.Request.Context(), uint(id), req.Valid)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "report resolved", report)
}
