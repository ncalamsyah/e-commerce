package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// GetDetailProduct
// @ID			GetDetailProduct
// @Tags		Product
// @Summary		Get Detail Product
// @Produce		json
// @Param		id	path	 integer	true "product ID"
// @Success		200 {object} response.SuccessResponse{data=entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/public/product/{id} [get]
func GetDetailProduct(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	res, err := repoProduct.GetDetailProduct(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
