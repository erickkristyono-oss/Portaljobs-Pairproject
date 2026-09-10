package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"admin-service/internal/dto/request"
	"admin-service/internal/usecase/user"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(
	userUsecase user.UserUsecase,
) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) UpdateStatus(
	c *echo.Context,
) error {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"responseCode":    "400",
				"responseMessage": "invalid user id",
				"responseData":    nil,
			},
		)
	}

	var req request.UpdateUserStatusRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"responseCode":    "400",
				"responseMessage": "invalid request body",
				"responseData":    nil,
			},
		)
	}

	if err := h.userUsecase.UpdateStatus(
		c.Request().Context(),
		uint(id),
		req,
	); err != nil {
		return c.JSON(
			http.StatusBadGateway,
			map[string]any{
				"responseCode":    "502",
				"responseMessage": err.Error(),
				"responseData":    nil,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]any{
			"responseCode":    "200",
			"responseMessage": "user status berhasil diperbarui",
			"responseData":    nil,
		},
	)
}
