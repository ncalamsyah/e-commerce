package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/product/dto"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
)

func UpdateProduct(id int, product *entity.Product) (dto.CreateProductResponse, error) {
	err := config.DB.
		Where("id = ?", id).
		Updates(&product).Error
	if err != nil {
		return dto.CreateProductResponse{}, err
	}
	return dto.CreateProductResponse{Id: uint(id), Message: "Success Update Product"}, nil
}
