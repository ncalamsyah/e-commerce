package helper

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ncalamsyah/e-commerce/models/response"
)

func (r *HttpResponse) SuccessResponse(ec echo.Context, message string, data interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusOK)
	}
	return ec.JSON(http.StatusOK, response.BuildSuccessResponse(message, http.StatusOK, data))
}

func (r *HttpResponse) SuccessResponseWithCode(ec echo.Context, code int, message string, data interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusOK)
	}
	return ec.JSON(http.StatusOK, response.BuildSuccessResponse(message, code, data))
}
