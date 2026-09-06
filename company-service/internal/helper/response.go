package helper

import "github.com/labstack/echo/v5"

type Response struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    interface{} `json:"responseData,omitempty"`
}

func Success(c *echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Response{
		ResponseCode:    statusCode,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func Error(c *echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, Response{
		ResponseCode:    statusCode,
		ResponseMessage: message,
	})
}

func BadRequest(c *echo.Context, message string) error {
	return Error(c, 400, message)
}

func Unauthorized(c *echo.Context, message string) error {
	return Error(c, 401, message)
}

func Forbidden(c *echo.Context, message string) error {
	return Error(c, 403, message)
}

func NotFound(c *echo.Context, message string) error {
	return Error(c, 404, message)
}

func Conflict(c *echo.Context, message string) error {
	return Error(c, 409, message)
}

func InternalServerError(c *echo.Context, message string) error {
	return Error(c, 500, message)
}
