package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
)

func GetOrderByUserID(id int) ([]entity.Transactions, error) {
	orders := []entity.Transactions{}
	err := config.DB.Where("customer_id = ?", id).
	Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
