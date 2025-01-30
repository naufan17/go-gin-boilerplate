package services

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/repositories"

	"errors"
)

func ProfileUser(id uuid.UUID) (dtos.ProfileDto, error) {
	user, err := repositories.GetUserByID(id)

	if err != nil {
		return dtos.ProfileDto{}, errors.New("not found")
	}

	return dtos.ToProfileDto(user), nil
}
