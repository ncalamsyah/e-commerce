package dto

type OrderRequestDTO struct {
	ProductID uint   `json:"product_id" form:"product_id" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required"`
}
