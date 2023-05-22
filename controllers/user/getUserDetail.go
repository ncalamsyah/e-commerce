package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	repository "github.com/ncalamsyah/e-commerce/repository/user"
)

// GetUserDetial
// @ID			GetUserDetail
// @Tags		User
// @Summary		GetUserDetail
// @Produce		json
// @Param		id	path	 integer	true	"user id"
// @Success		200	{object} response.SuccessResponse{data=entity.Users}
// @Failure		500	{object} response.ErrorResponse{errors=[]string}
// @Failure		400	{object} response.ErrorResponse{errors=[]string}
// @Router		/user/{id} [get]
func GetUserDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}
	logged := middlewares.ExtractTokenId(c)

	if id != logged {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error", http.StatusUnauthorized, errors.New("forbidden")))
	}

	user, err := repository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("user not found")))
	}

	if user.IsSeller {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("error", http.StatusUnauthorized, errors.New("user is not a customer")))
	}

	res, err := repository.GetUserDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("user not found")))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))

}
