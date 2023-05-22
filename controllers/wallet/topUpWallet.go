package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	order "github.com/ncalamsyah/e-commerce/models/order/entity"
	"github.com/ncalamsyah/e-commerce/models/response"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
	walletRepo "github.com/ncalamsyah/e-commerce/repository/wallet"
)

// TopUpWallet
// @ID			TopUpWallet
// @Tags		Wallet
// @Summary		Top-up Wallet
// @Security	JWTAuth
// @Produce		json
// @Param		amount		 formData	integer	true	"amount"
// @Success		200 {object} response.SuccessResponse{data=entity.Wallet}
// @Failure		500 {object} response.ErrorResponse{errors=[]string}
// @Failure		400 {object} response.ErrorResponse{errors=[]string}
// @Router		/user/wallet-topup [post]
func TopUpWallet(c echo.Context) error {
	req := entity.WalletTransactions{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BuildErrorResponse("error bind request", http.StatusBadRequest, err))
	}
	logged := middlewares.ExtractTokenId(c)

	wallet, err := walletRepo.GetWalletByUserId(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	expireTime := time.Now().Add(1 * time.Minute)
	data := entity.WalletTransactions{
		WalletID:   wallet.ID,
		CustomerID: uint(logged),
		Type:       entity.TopUp,
		Amount:     req.Amount,
		Status:     order.WaitingStatus,
		ExpiredAt:  expireTime,
	}

	res, err := walletRepo.WalletTrans(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("failed", http.StatusInternalServerError, err))
	}

	expirationTimer := time.NewTimer(1 * time.Minute)
	go func() {
		<-expirationTimer.C
		if data.Status == order.WaitingStatus {
			data.Status = order.ExpiredStatus
			_, err := walletRepo.UpdateWalletTrans(int(res.ID), data.Status)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
