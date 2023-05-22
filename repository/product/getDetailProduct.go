package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
)

func GetDetailProduct(id int) (entity.Product, error) {
	product := entity.Product{}
	err := config.DB.
		Where("id = ?", id).
		First(&product)
	if err.Error != nil {
		return entity.Product{}, err.Error
	}
	return product, nil
}
