package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/dto"
)

func UpdateWalletTrans(id int, status string) (dto.UpdateWalletTransResponse, error) {
	err := config.DB.Table("wallet_transactions").
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return dto.UpdateWalletTransResponse{}, err
	}

	return dto.UpdateWalletTransResponse{
		ID:      uint(id),
		Message: "success",
	}, nil
}
