package dtos

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

type ProfileDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UpdateProfileDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdatePasswordDto struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ToProfileDto(user models.User) ProfileDto {
	return ProfileDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
