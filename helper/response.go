package helper

import "net/http"

func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
	}
	return result
}

func UnauthorizedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": "Unauthorized",
	}
	return result
}

func SuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    data,
	}
	return result
}

func InternalServerResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Internal Server Error",
	}
	return result
}

func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Login Failed",
	}
	return result
}

func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    data,
	}
	return result
}
