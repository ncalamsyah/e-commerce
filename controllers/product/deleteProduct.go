package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// DeleteProduct
// @ID			DeleteProduct
// @Tags		Product
// @Summary		Delete Product
// @Security	JWTAuth
// @Produce		json
// @Param		id	path	 integer	true "product ID"
// @Success		200 {object} response.SuccessResponse{data=entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/product/{id} [delete]
func DeleteProduct(c echo.Context) error {
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
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("failed", http.StatusInternalServerError, errors.New("unable to delete product")))
	}

	res, err := repoProduct.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
