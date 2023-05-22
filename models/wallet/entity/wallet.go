package entity

import "github.com/ncalamsyah/e-commerce/models/response"

type Wallet struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Balance float64 `json:"balance"`

	response.BaseModelSoftDelete
}
