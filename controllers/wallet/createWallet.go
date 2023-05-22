package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
	userRepo "github.com/ncalamsyah/e-commerce/repository/user"
	walletRepo "github.com/ncalamsyah/e-commerce/repository/wallet"
)

// CreateWallet
// @ID			CreateWallet
// @Tags		Wallet
// @Summary		Create Wallet
// @Produce		json
// @Security	JWTAuth
// @Success		200	{object} response.SuccessResponse{data=dto.CreateWalletResponse}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/user/wallet [post]
func CreateWallet(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)

	user, err := userRepo.GetUser(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	if user.IsSeller {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, errors.New("seller unable to create wallet")))
	}
	
	data := entity.Wallet{
		CustomerID: uint(logged),
	}

	res, err := walletRepo.CreateWallet(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error create wallet", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusCreated, response.BuildSuccessResponse("success", http.StatusOK, res))
}
