package helper

import "github.com/labstack/echo/v5"

type Response struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    interface{} `json:"responseData,omitempty"`
}

func Success(c *echo.Context, message string, data interface{}) error {
	return c.JSON(200, Response{
		ResponseCode:    200,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func Created(c *echo.Context, message string, data interface{}) error {
	return c.JSON(201, Response{
		ResponseCode:    201,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func BadRequest(c *echo.Context, message string) error {
	return c.JSON(400, Response{
		ResponseCode:    400,
		ResponseMessage: message,
	})
}

func Unauthorized(c *echo.Context, message string) error {
	return c.JSON(401, Response{
		ResponseCode:    401,
		ResponseMessage: message,
	})
}

func Forbidden(c *echo.Context, message string) error {
	return c.JSON(403, Response{
		ResponseCode:    403,
		ResponseMessage: message,
	})
}

func NotFound(c *echo.Context, message string) error {
	return c.JSON(404, Response{
		ResponseCode:    404,
		ResponseMessage: message,
	})
}

func Conflict(c *echo.Context, message string) error {
	return c.JSON(409, Response{
		ResponseCode:    409,
		ResponseMessage: message,
	})
}

func InternalServerError(c *echo.Context, message string) error {
	return c.JSON(500, Response{
		ResponseCode:    500,
		ResponseMessage: message,
	})
}
