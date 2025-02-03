package services

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/repositories"
	"golang.org/x/crypto/bcrypt"

	"errors"
)

func ProfileUser(id uuid.UUID) (dtos.ProfileDto, error) {
	user, err := repositories.GetUserByID(id)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("not found")
	}

	return dtos.ToProfileDto(user), nil
}

func UpdateProfileUser(user dtos.UpdateProfileDto, id uuid.UUID) (dtos.ProfileDto, error) {
	userUpdate, err := repositories.UpdateProfile(user, id)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("not found")
	}

	return dtos.ToProfileDto(userUpdate), nil
}

func UpdatePasswordUser(user dtos.UpdatePasswordDto, id uuid.UUID) (dtos.ProfileDto, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	userUpdate, err := repositories.UpdatePassword(user, id)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("not found")
	}

	return dtos.ToProfileDto(userUpdate), nil
}
