package handler

import (
	domainerrors "company-service/internal/domain/error"
	"company-service/internal/helper"

	"github.com/labstack/echo/v5"
)

func handleDomainError(c *echo.Context, err error) error {
	switch err {
	case domainerrors.ErrNotFound:
		return helper.NotFound(c, err.Error())

	case domainerrors.ErrConflict:
		return helper.Conflict(c, err.Error())

	case domainerrors.ErrForbidden:
		return helper.Forbidden(c, err.Error())

	default:
		return helper.InternalServerError(c, "internal server error")
	}
}
