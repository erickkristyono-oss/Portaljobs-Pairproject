package handler

import (
	"net/http"

	"user-service/internal/dto/request"
	"user-service/internal/helper"
	"user-service/internal/middleware"
	profileUsecase "user-service/internal/usecase/profile"

	"github.com/labstack/echo/v5"
)

type ProfileHandler struct {
	profileUsecase profileUsecase.ProfileUsecase
}

func NewProfileHandler(
	profileUsecase profileUsecase.ProfileUsecase,
) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
	}
}

func getUserID(c *echo.Context) (uint, error) {
	userID, ok := c.Get(middleware.UserIDKey).(uint)

	if !ok {
		return 0, helper.Unauthorized(
			c,
			"invalid user identity",
		)
	}

	return userID, nil
}

func (h *ProfileHandler) Create(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	var req request.CreateProfileRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.profileUsecase.Create(
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
		"profile berhasil dibuat",
		result,
	)
}

func (h *ProfileHandler) GetByUserID(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	result, err := h.profileUsecase.GetByUserID(
		c.Request().Context(),
		userID,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"profile berhasil ditemukan",
		result,
	)
}

func (h *ProfileHandler) Update(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	var req request.UpdateProfileRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.profileUsecase.Update(
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
		"profile berhasil diperbarui",
		result,
	)
}
