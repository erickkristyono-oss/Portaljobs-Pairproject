package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"company-service/internal/dto/request"
	"company-service/internal/helper"
	"company-service/internal/usecase/company_profile"
)

type CompanyProfileHandler struct {
	usecase company_profile.Usecase
}

func NewCompanyProfileHandler(
	usecase company_profile.Usecase,
) *CompanyProfileHandler {
	return &CompanyProfileHandler{
		usecase: usecase,
	}
}

func (h *CompanyProfileHandler) Create(c *echo.Context) error {

	// Ambil user_id dari JWT
	userID, ok := c.Get("user_id").(uint)

	if !ok {
		return helper.Error(c, 401, "invalid user")
	}

	var req request.CreateCompanyProfileRequest

	// Bind request body
	if err := c.Bind(&req); err != nil {
		return helper.Error(c, 400, "invalid request body")
	}

	// userID berasal dari JWT
	result, err := h.usecase.Create(
		context.Background(),
		userID,
		&req,
	)

	if err != nil {
		return helper.Error(c, 400, err.Error())
	}

	return helper.Success(
		c,
		201,
		"company profile created successfully",
		result,
	)
}

func (h *CompanyProfileHandler) Get(c *echo.Context) error {

	userID, ok := c.Get("user_id").(uint)

	if !ok {
		return helper.Error(c, 401, "invalid user")
	}

	result, err := h.usecase.Get(
		context.Background(),
		userID,
	)

	if err != nil {
		return helper.Error(c, 404, err.Error())
	}

	return helper.Success(
		c,
		200,
		"company profile retrieved successfully",
		result,
	)
}

func (h *CompanyProfileHandler) Update(c *echo.Context) error {

	userID, ok := c.Get("user_id").(uint)

	if !ok {
		return helper.Error(c, 401, "invalid user")
	}

	var req request.UpdateCompanyProfileRequest

	if err := c.Bind(&req); err != nil {
		return helper.Error(c, 400, "invalid request body")
	}

	result, err := h.usecase.Update(
		context.Background(),
		userID,
		&req,
	)

	if err != nil {
		return helper.Error(c, 404, err.Error())
	}

	return helper.Success(
		c,
		200,
		"company profile updated successfully",
		result,
	)
}

func (h *CompanyProfileHandler) GetByUserID(c *echo.Context) error {

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

	result, err := h.usecase.Get(
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
		"company profile retrieved successfully",
		result,
	)
}
