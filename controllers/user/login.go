package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/models/response"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
	repository "github.com/ncalamsyah/e-commerce/repository/user"
)

// Login
// @ID			Login
// @Tags		User
// @Summary		Login User
// @Produce		json
// @Param		email		formData	string	true	"email"
// @Param		password	formData	string	true	"password"
// @Success		200	{object} response.SuccessResponse{data=[]string}
// @Failure		500	{object} response.ErrorResponse{errors=[]string}
// @Failure		400	{object} response.ErrorResponse{errors=[]string}
// @Router		/public/users/login [post]
func Login(c echo.Context) error {
	req := entity.Users{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}
	pass := req.Password
	res, err := repository.Login(pass, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("login failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
