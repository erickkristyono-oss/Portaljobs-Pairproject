package handler

import (
	"net/http"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct{ companies *usecase.CompanyUsecase }

func NewCompanyHandler(companies *usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{companies: companies}
}

func (h *CompanyHandler) Save(c *gin.Context) {
	var req CompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	company, err := h.companies.Save(c.Request.Context(), middleware.UserID(c),
		req.Name, req.FieldOf, req.Address, req.Description)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "company profile saved", company)
}

func (h *CompanyHandler) Get(c *gin.Context) {
	company, err := h.companies.Get(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "company profile", company)
}
