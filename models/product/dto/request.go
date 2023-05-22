package dto

type CreateProductRequestDTO struct {
	Name     string `json:"name" form:"name" validate:"required"`
	SellerID uint   `json:"seller_id" form:"-"`
	Quantity int    `json:"quantity" form:"quantity" validate:"required"`
	Price    string `json:"price" form:"price" validate:"required"`
}
