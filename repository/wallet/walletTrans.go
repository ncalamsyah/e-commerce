package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/wallet/entity"
)

func WalletTrans(data *entity.WalletTransactions) (entity.WalletTransactions, error) {
	err := config.DB.Create(&data).Error
	if err != nil {
		return entity.WalletTransactions{}, err
	}

	return entity.WalletTransactions{
		ID:                  data.ID,
		CustomerID:          data.CustomerID,
		WalletID:            data.WalletID,
		Type:                data.Type,
		Amount:              data.Amount,
		Status:              data.Status,
		ExpiredAt:           data.ExpiredAt,
		BaseModelSoftDelete: data.BaseModelSoftDelete,
	}, nil
}
