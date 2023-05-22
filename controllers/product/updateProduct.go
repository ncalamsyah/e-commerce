package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/product/dto"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
	"gopkg.in/go-playground/validator.v9"
)

// UpdateProduct
// @ID			UpdateProduct
// @Tags		Product
// @Summary		Update Products
// @Security	JWTAuth
// @Produce		json
// @Param		id			path		integer	true	"product id"
// @Param		name		formData	string	true	"product name"
// @Param		quantity	formData	integer	true	"quantity"
// @Param		price		formData	string	true	"price"
// @Success		200 {object} response.SuccessResponse{data=entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/product/{id} [put]
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	product, err := repoProduct.GetDetailProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	logged := middlewares.ExtractTokenId(c)

	if logged != int(product.SellerID) {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("failed", http.StatusInternalServerError, errors.New("unable to update product")))
	}

	req := dto.CreateProductRequestDTO{}
	c.Bind(&req)
	v := validator.New()
	e := v.Struct(req)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error validating request", http.StatusBadRequest, err))
	}

	data := entity.Product{
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	res, err := repoProduct.UpdateProduct(id, &data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
