package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/product/dto"
	"github.com/ncalamsyah/e-commerce/models/product/entity"
)

func DeleteProduct(id int) (dto.CreateProductResponse, error) {
	product := entity.Product{}
	checkProduct := config.DB.Find(&product, id).RowsAffected
	err := config.DB.Delete(&product).Error
	if err != nil || checkProduct > 0 {
		return dto.CreateProductResponse{Id: uint(id), Message: "Product Deleted"}, err
	}
	return dto.CreateProductResponse{}, nil
}
