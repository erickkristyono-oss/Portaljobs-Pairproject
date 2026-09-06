package helper

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Response struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    interface{} `json:"responseData,omitempty"`
}

func Success(c *echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func Created(c *echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusCreated, Response{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func BadRequest(c *echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, Response{
		ResponseCode:    http.StatusBadRequest,
		ResponseMessage: message,
	})
}

func Unauthorized(c *echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, Response{
		ResponseCode:    http.StatusUnauthorized,
		ResponseMessage: message,
	})
}

func Forbidden(c *echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, Response{
		ResponseCode:    http.StatusForbidden,
		ResponseMessage: message,
	})
}

func NotFound(c *echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, Response{
		ResponseCode:    http.StatusNotFound,
		ResponseMessage: message,
	})
}

func Conflict(c *echo.Context, message string) error {
	return c.JSON(http.StatusConflict, Response{
		ResponseCode:    http.StatusConflict,
		ResponseMessage: message,
	})
}

func InternalServerError(c *echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, Response{
		ResponseCode:    http.StatusInternalServerError,
		ResponseMessage: message,
	})
}
