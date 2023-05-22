package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
)

func GetWalletTransById(id int) (entity.WalletTransactions, error) {
	trans := entity.WalletTransactions{}
	err := config.DB.Where("id = ?", id).
		First(&trans).Error
	if err != nil {
		return entity.WalletTransactions{}, err
	}
	return trans, nil
}
