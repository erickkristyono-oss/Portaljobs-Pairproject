package handler

import (
	"net/http"
	"strconv"

	domainerrors "user-service/internal/domain/error"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/helper"
	userUsecase "user-service/internal/usecase/user"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	userUsecase userUsecase.UserUsecase
}

func NewUserHandler(
	userUsecase userUsecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// Register godoc
func (h *UserHandler) Register(c *echo.Context) error {
	var req request.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.userUsecase.Register(
		c.Request().Context(),
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusCreated,
		"register berhasil",
		result,
	)
}

// Login godoc
func (h *UserHandler) Login(c *echo.Context) error {
	var req request.LoginRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	user, err := h.userUsecase.Login(
		c.Request().Context(),
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	token, err := helper.GenerateToken(
		user.ID,
		string(user.Role),
	)
	if err != nil {
		return helper.InternalServerError(
			c,
			"failed to generate token",
		)
	}

	result := response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        user.ID,
			Nama:      user.Nama,
			Email:     user.Email,
			Role:      string(user.Role),
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return helper.Success(
		c,
		http.StatusOK,
		"login berhasil",
		result,
	)
}

// GetByID godoc
func (h *UserHandler) GetByID(c *echo.Context) error {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid user id",
		)
	}

	result, err := h.userUsecase.GetByID(
		c.Request().Context(),
		uint(id),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"user berhasil ditemukan",
		result,
	)
}

// GetAll godoc
func (h *UserHandler) GetAll(c *echo.Context) error {
	result, err := h.userUsecase.GetAll(
		c.Request().Context(),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"data user berhasil ditemukan",
		result,
	)
}

// UpdateStatus godoc
func (h *UserHandler) UpdateStatus(c *echo.Context) error {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid user id",
		)
	}

	var req request.UpdateUserStatusRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	if err := h.userUsecase.UpdateStatus(
		c.Request().Context(),
		uint(id),
		req,
	); err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"user status berhasil diperbarui",
		nil,
	)
}

func handleDomainError(
	c *echo.Context,
	err error,
) error {
	switch err {
	case domainerrors.ErrNotFound:
		return helper.NotFound(
			c,
			err.Error(),
		)

	case domainerrors.ErrEmailTaken,
		domainerrors.ErrConflict:
		return helper.Conflict(
			c,
			err.Error(),
		)

	case domainerrors.ErrInvalidCredential,
		domainerrors.ErrSuspended:
		return helper.Unauthorized(
			c,
			err.Error(),
		)

	case domainerrors.ErrForbidden:
		return helper.Forbidden(
			c,
			err.Error(),
		)

	case domainerrors.ErrInvalidRole,
		domainerrors.ErrInvalidLevel,
		domainerrors.ErrInvalidStatus:
		return helper.BadRequest(
			c,
			err.Error(),
		)

	default:
		return helper.InternalServerError(
			c,
			"internal server error",
		)
	}
}
