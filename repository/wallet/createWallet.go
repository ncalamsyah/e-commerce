package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/dto"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
)

func CreateWallet(data *entity.Wallet) (dto.CreateWalletResponse, error) {
	if err := config.DB.Create(&data).Error; err != nil {
		return dto.CreateWalletResponse{}, err
	}
	return dto.CreateWalletResponse{
		ID:      data.ID,
		Message: "Success Create Wallet",
	}, nil
}
