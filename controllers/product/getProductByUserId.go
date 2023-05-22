package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// GetProductByUserID
// @ID			GetProductByUserId
// @Tags		Product
// @Summary		Get Product By User ID
// @Produce		json
// @Param		user_id		 path	 integer	true "user ID"
// @Success		200 {object} response.SuccessResponse{data=[]entity.Product}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/public/products/{user_id} [get]
func GetProductByUserId(c echo.Context) error {
	id := c.Param("user_id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	res, err := repoProduct.GetProductByUserId(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
