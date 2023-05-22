package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
)

func CreateOrder(order *entity.Transactions) (entity.Transactions, error) {
	err := config.DB.Create(&order).Error
	if err != nil {
		return entity.Transactions{}, err
	}

	return entity.Transactions{
		ID:                  order.ID,
		CustomerID:          order.CustomerID,
		ProductID:           order.ProductID,
		Quantity:            order.Quantity,
		TotalPrice:          order.TotalPrice,
		Address:             order.Address,
		Status:              order.Status,
		ExpiredAt:           order.ExpiredAt,
		BaseModelSoftDelete: order.BaseModelSoftDelete,
	}, nil
}
