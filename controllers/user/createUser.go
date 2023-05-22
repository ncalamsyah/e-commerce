package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/helper"
	"github.com/ncalamsyah/e-commerce/models/response"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
	repository "github.com/ncalamsyah/e-commerce/repository/user"
	"gopkg.in/go-playground/validator.v9"
)

// CreateUser
// @ID			CreateUser
// @Tags		User
// @Summary		Register New User
// @Produce		json
// @Param		name		formData	string	true	"name"
// @Param		email		formData	string	true	"email"
// @Param		password	formData	string	true	"password"
// @Success		200 {object} response.SuccessResponse{data=dto.CreateUserResponse}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/public/users/register [post]
func CreateUser(c echo.Context) error {
	req := dto.UserRequestDTO{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}

	v := validator.New()
	err := v.Struct(req)
	if err == nil {
		req.Password, _ = helper.HashPassword(req.Password)
	}

	data := entity.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := repository.CreateUser(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error create user", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusCreated, response.BuildSuccessResponse("success", http.StatusOK, res))
}
