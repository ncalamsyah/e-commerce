package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
)

func GetUser(id int) (dto.GetUserPublicResponse, error) {
	users := entity.Users{}
	err := config.DB.Find(&users, id)
	rowsAffected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rowsAffected < 1 {
		return dto.GetUserPublicResponse{}, err.Error
	}
	return dto.GetUserPublicResponse{
		Name:     users.Name,
		Email:    users.Email,
		IsSeller: users.IsSeller,
	}, nil
}
