package handler

import (
	"errors"

	domainerrors "job-service/internal/domain/error"
	"job-service/internal/helper"

	"github.com/labstack/echo/v5"
)

func handleDomainError(c *echo.Context, err error) error {
	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		return helper.NotFound(c, "job not found")

	case errors.Is(err, domainerrors.ErrForbidden):
		return helper.Forbidden(c, "forbidden")

	case errors.Is(err, domainerrors.ErrConflict):
		return helper.Conflict(c, "job conflict")

	default:
		return helper.InternalServerError(c, "internal server error")
	}
}
