package repositories

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
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

func UpdateProfile(user dtos.UpdateProfileDto, id uuid.UUID) (models.User, error) {
	var updatedUser models.User

	if err := configs.DB.Model(&updatedUser).Where("id = ?", id).Updates(models.User{Name: user.Name, Email: user.Email}).Error; err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func UpdatePassword(user dtos.UpdatePasswordDto, id uuid.UUID) (models.User, error) {
	var updatedUser models.User

	if err := configs.DB.Model(&updatedUser).Where("id = ?", id).Updates(models.User{Password: user.Password}).Error; err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}
