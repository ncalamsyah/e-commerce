package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
)

func GetOrderDetail(id int) (entity.Transactions, error) {
	order := entity.Transactions{}
	err := config.DB.Where("id = ?", id).
		First(&order).Error
	if err != nil {
		return entity.Transactions{}, err
	}
	return order, nil
}
