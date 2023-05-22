package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
)

func GetOrderByProductIds(productIds []int) ([]entity.Transactions, error) {
	orders := []entity.Transactions{}
	err := config.DB.Where("product_id IN ?", productIds).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
