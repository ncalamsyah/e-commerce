package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/product/dto"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
)

func CreateProduct(product *entity.Product) (dto.CreateProductResponse, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return dto.CreateProductResponse{}, err
	}
	return dto.CreateProductResponse{Id: product.ID, Message: "Success Create Product"}, nil
}
