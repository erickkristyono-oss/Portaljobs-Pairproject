package helper

import "github.com/labstack/echo/v5"

type Response struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ResponseData    any    `json:"responseData"`
}

func Success(
	c *echo.Context,
	code int,
	message string,
	data any,
) error {
	return c.JSON(code, Response{
		ResponseCode:    "00",
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func Error(
	c *echo.Context,
	code int,
	message string,
) error {
	return c.JSON(code, Response{
		ResponseCode:    "ERROR",
		ResponseMessage: message,
		ResponseData:    nil,
	})
}
