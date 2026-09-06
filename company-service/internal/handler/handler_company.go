package handler

import (
	"net/http"
	"strconv"

	"company-service/internal/dto/request"
	"company-service/internal/helper"
	"company-service/internal/middleware"
	companyUsecase "company-service/internal/usecase/company"

	"github.com/labstack/echo/v5"
)

type CompanyHandler struct {
	companyUsecase companyUsecase.CompanyUsecase
}

func NewCompanyHandler(
	companyUsecase companyUsecase.CompanyUsecase,
) *CompanyHandler {
	return &CompanyHandler{
		companyUsecase: companyUsecase,
	}
}

// Create membuat company profile untuk user yang sedang login.
func (h *CompanyHandler) Create(c *echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return helper.Unauthorized(c, "invalid user identity")
	}

	var req request.CreateCompanyRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, "invalid request body")
	}

	result, err := h.companyUsecase.Create(
		c.Request().Context(),
		userID,
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusCreated,
		"company berhasil dibuat",
		result,
	)
}

// GetByID mengambil company berdasarkan ID.
func (h *CompanyHandler) GetByID(c *echo.Context) error {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil || id == 0 {
		return helper.BadRequest(c, "invalid company id")
	}

	result, err := h.companyUsecase.GetByID(
		c.Request().Context(),
		uint(id),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company berhasil ditemukan",
		result,
	)
}

// GetByUserID mengambil company milik user yang sedang login.
func (h *CompanyHandler) GetByUserID(c *echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return helper.Unauthorized(c, "invalid user identity")
	}

	result, err := h.companyUsecase.GetByUserID(
		c.Request().Context(),
		userID,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company berhasil ditemukan",
		result,
	)
}

// Update mengubah company milik user yang sedang login.
func (h *CompanyHandler) Update(c *echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return helper.Unauthorized(c, "invalid user identity")
	}

	var req request.UpdateCompanyRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, "invalid request body")
	}

	result, err := h.companyUsecase.Update(
		c.Request().Context(),
		userID,
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company berhasil diperbarui",
		result,
	)
}

// GetCompanyByUserID mengambil company berdasarkan user ID.
// Endpoint ini digunakan untuk komunikasi internal antar service.
func (h *CompanyHandler) GetCompanyByUserID(c *echo.Context) error {
	userID, err := strconv.ParseUint(
		c.Param("user_id"),
		10,
		64,
	)
	if err != nil || userID == 0 {
		return helper.BadRequest(c, "invalid user id")
	}

	result, err := h.companyUsecase.GetByUserID(
		c.Request().Context(),
		uint(userID),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company berhasil ditemukan",
		result,
	)
}
