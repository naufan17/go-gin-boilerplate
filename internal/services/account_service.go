package services

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
	"github.com/naufan17/go-gin-boilerplate/internal/repositories"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"

	"errors"
)

func ProfileUser(id uuid.UUID) (dtos.ProfileDto, error) {
	user, err := repositories.GetUserByID(id)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("not found")
	}

	return dtos.ToProfileDto(user), nil
}

func SessionUser(id uuid.UUID) ([]dtos.SessionDto, error) {
	session, err := repositories.GetSessionByUserID(id)

	if err != nil {
		return []dtos.SessionDto{}, errors.New("not found")
	}

	var sessionDtos []dtos.SessionDto
	for _, s := range session {
		sessionDtos = append(sessionDtos, dtos.ToSessionDto(s))
	}

	return sessionDtos, nil
}

func UpdateProfileUser(user dtos.UpdateProfileDto, id uuid.UUID) (models.User, error) {
	userUpdate, err := repositories.UpdateProfile(user, id)

	if err != nil {
		return models.User{}, errors.New("not found")
	}

	return userUpdate, nil
}

func UpdatePasswordUser(user dtos.UpdatePasswordDto, id uuid.UUID) (models.User, error) {
	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		return models.User{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	userUpdate, err := repositories.UpdatePassword(user, id)

	if err != nil {
		return models.User{}, errors.New("not found")
	}

	return userUpdate, nil
}
