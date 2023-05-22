package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoOrder "github.com/ncalamsyah/e-commerce/repository/order"
	repoProduct "github.com/ncalamsyah/e-commerce/repository/product"
	repoUser "github.com/ncalamsyah/e-commerce/repository/user"
)

// GetOrderByUserIdOrProductIds
// @ID			GetOrderByUserIdOrProductIds
// @Tags		Order
// @Summary		Get Order By User ID or Product IDs
// @Security	JWTAuth
// @Produce		json
// @Param		user_id	path	 integer	true "user id"
// @Success		200 {object} response.SuccessResponse{data=[]entity.Transactions}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/order/order-list [get]
func GetOrderByUserIdOrProductIds(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)

	user, err := repoUser.GetUser(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error get user", http.StatusInternalServerError, err))
	}

	if user.IsSeller {
		var productIds []int
		products, err := repoProduct.GetProductByUserId(logged)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
		}
		for _, product := range products {
			productIds = append(productIds, int(product.ID))
		}

		res, err := repoOrder.GetOrderByProductIds(productIds)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
		}

		return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
	}

	res, err := repoOrder.GetOrderByUserID(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
