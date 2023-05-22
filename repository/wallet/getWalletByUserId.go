package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
)

func GetWalletByUserId(id int) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	err := config.DB.
		Where("customer_id = ?", id).
		Find(&wallet)
	rowsAffected := config.DB.Where("customer_id = ?", id).Find(&wallet).RowsAffected
	if err.Error != nil || rowsAffected < 1 {
		return entity.Wallet{}, err.Error
	}
	return entity.Wallet{
		ID:                  wallet.ID,
		CustomerID:          wallet.CustomerID,
		Balance:             wallet.Balance,
		BaseModelSoftDelete: wallet.BaseModelSoftDelete,
	}, nil
}
