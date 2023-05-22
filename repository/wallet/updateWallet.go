package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/dto"
)

func UpdateWallet(id int, amount int) (dto.UpdateWalletTransResponse, error) {
	err := config.DB.Table("wallets").
		Where("id = ?", id).
		Update("balance", amount).Error
	if err != nil {
		return dto.UpdateWalletTransResponse{}, err
	}
	return dto.UpdateWalletTransResponse{ID: uint(id), Message: "success"}, nil
}
