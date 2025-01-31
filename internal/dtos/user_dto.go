package dtos

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

type RegisterDto struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=10"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" min:"10"`
}

type ProfileDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type AccessTokenDto struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func ToUserModel(user RegisterDto) models.User {
	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password}
}

func ToProfileDto(user models.User) ProfileDto {
	return ProfileDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
