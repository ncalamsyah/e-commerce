package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoOrder "github.com/ncalamsyah/e-commerce/repository/order"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// GetOrderDetail
// @ID			GetOrderDetail
// @Tags		Order
// @Summary		Get Order Detail
// @Security	JWTAuth
// @Produce		json
// @Param		id	path	 integer	true "order id"
// @Success		200 {object} response.SuccessResponse{data=entity.Transactions}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/order/{id} [get]
func GetOrderDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	logged := middlewares.ExtractTokenId(c)

	res, err := repoOrder.GetOrderDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	product, err := repoProduct.GetDetailProduct(int(res.ProductID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	if res.CustomerID != uint(logged) && product.SellerID != uint(logged) {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("err", http.StatusBadRequest, errors.New("unable to get order detail")))
	}

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
