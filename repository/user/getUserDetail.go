package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
)

func GetUserDetail(id int) (dto.GetUserDetailResponse, error) {
	user := entity.Users{}
	err := config.DB.Find(&user, id)
	rowsAffected := config.DB.Find(&user, id).RowsAffected
	if err.Error != nil || rowsAffected < 1 {
		return dto.GetUserDetailResponse{}, err.Error
	}
	return dto.GetUserDetailResponse{
		Name:     user.Name,
		Email:    user.Email,
		IsSeller: user.IsSeller,
	}, nil
}
