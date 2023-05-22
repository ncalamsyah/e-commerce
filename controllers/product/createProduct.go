package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/product/dto"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
	repoUser "github.com/ncalamsyah/e-commerce/repository/user"
	"gopkg.in/go-playground/validator.v9"
)

// CreateProduct
// @ID			CreateProduct
// @Tags		Product
// @Summary		Create Products
// @Security	JWTAuth
// @Produce		json
// @Param		name		formData	string	true	"product name"
// @Param		quantity	formData	integer	true	"quantity"
// @Param		price		formData	string	true	"price"
// @Success		200 {object} response.SuccessResponse{data=entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/product [post]
func CreateProduct(c echo.Context) error {
	req := dto.CreateProductRequestDTO{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}

	v := validator.New()
	e := v.Struct(req)
	if e == nil {
		logged := middlewares.ExtractTokenId(c)
		req.SellerID = uint(logged)
	}

	user, err := repoUser.GetUser(int(req.SellerID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error get user", http.StatusInternalServerError, err))
	}

	if !user.IsSeller {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("error", http.StatusUnauthorized, errors.New("user is not a seller")))
	}
	data := entity.Product{
		Name:     req.Name,
		SellerID: req.SellerID,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	res, err := repoProduct.CreateProduct(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusCreated, response.BuildSuccessResponse("success", http.StatusCreated, res))
}
