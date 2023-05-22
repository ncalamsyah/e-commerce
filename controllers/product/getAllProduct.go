package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// GetAllProducts
// @ID			GetAllProducts
// @Tags		Product
// @Summary		Get All Products
// @Produce		json
// @Success		200 {object} response.SuccessResponse{data=[]entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/public/products [get]
func GetAllProduct(c echo.Context) error {
	res, err := repoProduct.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
