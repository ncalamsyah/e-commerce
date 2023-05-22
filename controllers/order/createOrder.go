package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/order/dto"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoOrder "github.com/ncalamsyah/e-commerce/repository/order"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
	repoUser "github.com/ncalamsyah/e-commerce/repository/user"
)

// CreateOrder
// @ID			CreateOrder
// @Tags		Order
// @Summary		Create Order
// @Security	JWTAuth
// @Produce		json
// @Param		product_id	 formData	integer		true	"product id"
// @Param		address		 formData	string		true	"address"
// @Param		quantity	 formData	integer		true	"quantity"
// @Success		200 {object} response.SuccessResponse{data=entity.Transactions}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/order [post]
func CreateOrder(c echo.Context) error {
	req := dto.OrderRequestDTO{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}
	logged := middlewares.ExtractTokenId(c)

	user, err := repoUser.GetUser(int(logged))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error get user", http.StatusInternalServerError, err))
	}

	if user.IsSeller {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("error", http.StatusUnauthorized, errors.New("user is not a customer")))
	}

	product, err := repoProduct.GetDetailProduct(int(req.ProductID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, err))
	}

	if req.Quantity > product.Quantity {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("quantity exceed product's stock")))
	}

	convPrice, err := strconv.Atoi(product.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, err))
	}

	totalPrice := req.Quantity * convPrice
	expireTime := time.Now().Add(1 * time.Minute)

	data := entity.Transactions{
		CustomerID: uint(logged),
		ProductID:  req.ProductID,
		Address:    req.Address,
		Quantity:   req.Quantity,
		TotalPrice: float64(totalPrice),
		Status:     entity.WaitingStatus,
		ExpiredAt:  expireTime,
	}

	res, err := repoOrder.CreateOrder(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	expirationTimer := time.NewTimer(1 * time.Minute)
	go func() {
		<-expirationTimer.C
		if data.Status == entity.WaitingStatus {
			data.Status = entity.ExpiredStatus
			_, err := repoOrder.UpdateOrder(int(res.ID), data.Status)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	return c.JSON(http.StatusCreated, response.BuildSuccessResponse("success", http.StatusCreated, res))
}
