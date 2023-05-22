package helper

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ncalamsyah/e-commerce/models/response"
	"go.uber.org/zap"
)

var (
	ErrBadRequest     = NewError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, nil)
	ErrUnauthorized   = NewError(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, nil)
	ErrNotFound       = NewError(http.StatusText(http.StatusNotFound), http.StatusNotFound, nil)
	ErrInternalServer = NewError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil)
	ErrForbidden      = NewError(http.StatusText(http.StatusForbidden), http.StatusForbidden, nil)
)

type HttpError struct {
	response.ErrorResponse
	err error
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewError(message string, status int, err error) *HttpError {
	if message == "" {
		message = http.StatusText(status)
	}
	if err == nil {
		err = errors.New("")
	}
	return &HttpError{
		ErrorResponse: response.BuildErrorResponse(message, status, err),
		err:           err,
	}
}

func ErrorWrap(base *HttpError, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	return NewError(base.Message, base.Code, err)
}

func ErrorWithMessage(base *HttpError, message string, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	return NewError(message, base.Code, err)
}

func ErrorWithErrMessage(base *HttpError, err error) *HttpError {
	if base == nil {
		base = ErrInternalServer
	}
	message := ""
	if err != nil {
		message = err.Error()
	}
	return NewError(message, base.Code, err)
}

func (r *HttpResponse) ErrorResponse(ec echo.Context, err error, request ...interface{}) error {
	httpErr, ok := err.(*HttpError)
	if !ok {
		httpErr = ErrorWrap(ErrInternalServer, err)
	}

	var req interface{}
	if len(request) > 0 {
		req = request[0]
	}

	r.logger.Error("", zap.Error(httpErr.err),
		zap.String("method", ec.Request().Method),
		zap.String("uri", ec.Request().RequestURI),
		zap.Any("request", req))

	return ec.JSON(httpErr.Code, response.BuildErrorResponse(httpErr.Message, httpErr.Code, httpErr.err))
}
