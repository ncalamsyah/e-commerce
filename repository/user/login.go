package repository

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/middlewares"
	"github.com/ncalamsyah/e-commerce/models/users/dto"
	"github.com/ncalamsyah/e-commerce/models/users/entity"
	"golang.org/x/crypto/bcrypt"
)

func Login(pass string, user *entity.Users) (*dto.LoginResponse, error) {
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if match != nil {
		return nil, match
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: user.Token}, nil
}
