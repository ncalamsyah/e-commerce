package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoOrder "github.com/ncalamsyah/e-commerce/repository/order"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
)

// UpdateOrder
// @ID			UpdateOrder
// @Tags		Order
// @Summary		Update Order Status
// @Produce		json
// @Param		id	path	 integer	true "order ID"
// @Success		200 {object} response.SuccessResponse{data=entity.Transactions}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/order/{id} [put]
func UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	order, err := repoOrder.GetOrderDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	product, err := repoProduct.GetDetailProduct(int(order.ProductID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	logged := middlewares.ExtractTokenId(c)

	if logged != int(product.SellerID) {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("failed", http.StatusInternalServerError, errors.New("unable to update order")))
	}

	var newStatus string
	switch order.Status {
	case entity.WaitingStatus:
		newStatus = entity.OnProcessStatus

	case entity.OnProcessStatus:
		newStatus = entity.ShippingProcess

	case entity.ShippingProcess:
		newStatus = entity.DeliveredStatus

	case entity.DeliveredStatus:
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("failed", http.StatusBadRequest, errors.New("order was delivered")))

	default:
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("failed", http.StatusBadRequest, errors.New("order is expired")))
	}

	res, err := repoOrder.UpdateOrder(id, newStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
