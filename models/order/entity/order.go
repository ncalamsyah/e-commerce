package entity

import (
	"time"

	"github.com/ncalamsyah/e-commerce/models/response"
)

const (
	WaitingStatus   = "waiting"
	OnProcessStatus = "on-process"
	ShippingProcess = "shipping"
	DeliveredStatus = "delivered"
	ExpiredStatus   = "expired"
)

type Transactions struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CustomerID uint      `json:"customer_id"`
	ProductID  uint      `json:"product_id"`
	Address    string    `json:"address"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	ExpiredAt  time.Time `json:"expired_at"`

	response.BaseModelSoftDelete
}
