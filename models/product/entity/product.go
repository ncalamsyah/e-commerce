package entity

import (
	"github.com/ncalamsyah/e-commerce/models/response"
)

type (
	Product struct {
		ID       uint   `json:"id" gorm:"primaryKey"`
		Name     string `json:"name"`
		SellerID uint   `json:"seller_id"`
		Quantity int    `json:"quantity"`
		Price    string `json:"price"`

		response.BaseModelSoftDelete
	}
)
