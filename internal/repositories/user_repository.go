package repositories

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

func CreateUser(user models.User) (models.User, error) {
	if err := configs.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if err := configs.DB.Where("email = ?", email).Select("id", "email", "password").First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := configs.DB.Where("id = ?", id).Select("id", "name", "email").First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
