package entity

import (
	"github.com/ncalamsyah/e-commerce/models/response"
)

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique" form:"email"`
	Password string `json:"password" form:"password"`
	IsSeller bool   `json:"is_seller"`
	WalletID uint   `json:"wallet_id"`
	Token    string

	response.BaseModelSoftDelete
}
