package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
)

func GetAllProduct() ([]entity.Product, error) {
	products := []entity.Product{}
	err := config.DB.Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}
