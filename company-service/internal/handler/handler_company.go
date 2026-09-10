package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"company-service/internal/helper"
	"company-service/internal/usecase/company"
)

type CompanyHandler struct {
	usecase company.Usecase
}

func NewCompanyHandler(usecase company.Usecase) *CompanyHandler {
	return &CompanyHandler{
		usecase: usecase,
	}
}

func (h *CompanyHandler) Create(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return helper.Error(c, http.StatusUnauthorized, "invalid user")
	}

	result, err := h.usecase.Create(
		context.Background(),
		userID,
	)
	if err != nil {
		return helper.Error(c, http.StatusConflict, err.Error())
	}

	return helper.Success(
		c,
		http.StatusCreated,
		"company created successfully",
		result,
	)
}

func (h *CompanyHandler) GetByUserID(c *echo.Context) error {
	userID, err := strconv.ParseUint(
		c.Param("user_id"),
		10,
		64,
	)
	if err != nil {
		return helper.Error(
			c,
			http.StatusBadRequest,
			"invalid user id",
		)
	}

	result, err := h.usecase.GetByUserID(
		c.Request().Context(),
		uint(userID),
	)
	if err != nil {
		return helper.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company retrieved successfully",
		result,
	)
}

func (h *CompanyHandler) GetByID(c *echo.Context) error {
	companyID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.Error(
			c,
			http.StatusBadRequest,
			"invalid company id",
		)
	}

	result, err := h.usecase.GetByID(
		c.Request().Context(),
		uint(companyID),
	)
	if err != nil {
		return helper.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"company retrieved successfully",
		result,
	)
}
