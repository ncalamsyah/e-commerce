package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/response"
	repoWallet "github.com/ncalamsyah/e-commerce/repository/wallet"
)

// GetWallet
// @ID			GetWallet
// @Tags		Wallet
// @Summary		Get Wallet
// @Security	JWTAuth
// @Produce		json
// @Success		200	{object} response.SuccessResponse{data=entity.Wallet}
// @Failure		500	{object} response.ErrorResponse{errors=[]string}
// @Failure		400	{object} response.ErrorResponse{errors=[]string}
// @Router		/user/wallet [get]
func GetWalletByUserId(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)

	res, err := repoWallet.GetWalletByUserId(logged)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("wallet not found")))
	}

	if res.ID == 0 {
		return c.JSON(http.StatusInternalServerError, response.BuildErrorResponse("error", http.StatusInternalServerError, errors.New("wallet not found")))
	}
	return c.JSON(http.StatusOK, response.BuildSuccessResponse("success", http.StatusOK, res))
}
