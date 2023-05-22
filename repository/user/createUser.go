package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
)

func CreateUser(data *entity.Users) (dto.CreateUserResponse, error) {
	if err := config.DB.Create(&data).Error; err != nil {
		return dto.CreateUserResponse{}, err
	}
	return dto.CreateUserResponse{Id: data.ID, Message: "User Register Success"}, nil
}
