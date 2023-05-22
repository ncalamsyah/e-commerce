package response

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type BaseModelSoftDelete struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func BuildSuccessResponse(message string, code int, data interface{}) SuccessResponse {
	res := SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, code int, err error) ErrorResponse {
	errorMessage := err.Error()

	splitError := strings.Split(errorMessage, "\n")
	res := ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  splitError,
	}
	return res
}
