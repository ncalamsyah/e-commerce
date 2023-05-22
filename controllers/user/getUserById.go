package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/models/response"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	repository "github.com/ncalamsyah/e-commerce/repository/user"
)

// GetUserByID
// @ID			GetUserByID
// @Tags		User
// @Summary		GetUserByID
// @Produce		json
// @Param		id	path	integer	true	"id"
// @Success		200	{object} response.SuccessResponse{data=entity.Users}
// @Failure		500	{object} response.ErrorResponse{errors=[]string}
// @Failure		400	{object} response.ErrorResponse{errors=[]string}
// @Router		/public/users/{id} [get]
func GetUser(c echo.Context) error {
	id := c.Param("id")
	nilRes := dto.GetUserResponse{}
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}
	res, err := repository.GetUser(convId)
	if err != nil || res == nilRes {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("user not found")))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))

}
