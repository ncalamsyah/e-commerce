package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoWallet "github.com/ncalamsyah/e-commerce/repository/wallet"
)

// ConfirmWalletTrans
// @ID			ConfirmWalletTrans
// @Tags		Wallet
// @Summary		Confirm Wallet Transaction
// @Produce		json
// @Param		id	path	 integer	true "wallet transaction ID"
// @Success		200 {object} response.SuccessResponse{data=entity.WalletTransactions}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/user/wallet-confirm/{id} [put]
func ConfirmWalletTrans(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error parsing id", http.StatusBadRequest, err))
	}

	trans, err := repoWallet.GetWalletTransById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	logged := middlewares.ExtractTokenId(c)

	if logged != int(trans.CustomerID) {
		return c.JSON(http.StatusUnauthorized, response.BuildErrorResponse("failed", http.StatusInternalServerError, errors.New("unable to cofirm wallet transaction")))
	}

	var newStatus string
	switch trans.Status {
	case entity.WaitingStatus:
		newStatus = "success"
	default:
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("failed", http.StatusBadRequest, errors.New("order is expired")))
	}

	res, err := repoWallet.UpdateWalletTrans(id, newStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	wallet, err := repoWallet.GetWalletByUserId(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	totalBalance := trans.Amount + wallet.Balance
	_, err = repoWallet.UpdateWallet(int(wallet.ID), totalBalance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
