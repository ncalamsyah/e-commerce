package entity

import (
	"time"

	"github.com/ncalamsyah/e-commerce/models/response"
)

const (
	TopUp    = "topup"
	WithDraw = "withdraw"
)

type Wallet struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CustomerID uint `json:"customer_id" gorm:"unique"`
	Balance    int  `json:"balance"`

	response.BaseModelSoftDelete
}

type WalletTransactions struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CustomerID uint      `json:"customer_id"`
	WalletID   uint      `json:"wallet_id"`
	Type       string    `json:"type"`
	Amount     int       `json:"amount" form:"amount" validate:"required"`
	Status     string    `json:"status"`
	ExpiredAt  time.Time `json:"expired_at"`

	response.BaseModelSoftDelete
}
