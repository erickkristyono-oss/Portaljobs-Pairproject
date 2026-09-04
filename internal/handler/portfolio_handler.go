package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/domain"
	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct{ portfolios *usecase.PortfolioUsecase }

func NewPortfolioHandler(p *usecase.PortfolioUsecase) *PortfolioHandler {
	return &PortfolioHandler{portfolios: p}
}

func (h *PortfolioHandler) Add(c *gin.Context) {
	var req PortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p := domain.Portfolio{NameProject: req.NameProject, Organization: req.Organization, Description: req.Description}
	saved, err := h.portfolios.Add(c.Request.Context(), middleware.UserID(c), p)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "portfolio added", saved)
}

func (h *PortfolioHandler) List(c *gin.Context) {
	items, err := h.portfolios.List(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "portfolios", items)
}

func (h *PortfolioHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid portfolio id")
		return
	}
	if err := h.portfolios.Remove(c.Request.Context(), middleware.UserID(c), uint(id)); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "portfolio deleted", nil)
}
