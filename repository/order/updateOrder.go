package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/order/dto"
	"github.com/ncalamsyah/e-commerce/models/order/entity"
)

func UpdateOrder(id int, status string) (dto.UpdateOrderResponse, error) {
	err := config.DB.Table("transactions").
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return dto.UpdateOrderResponse{}, err
	}

	var resMessage string
	switch status {
	case entity.OnProcessStatus:
		resMessage = "order on-process"

	case entity.ShippingProcess:
		resMessage = "order on shipping"

	case entity.DeliveredStatus:
		resMessage = "order delivered"

	default:
		resMessage = "failed update order"
	}
	return dto.UpdateOrderResponse{ID: uint(id), Message: resMessage}, nil
}
