package dtos

import (
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

type RegisterDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required" min:"10"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required" min:"10"`
}

type ProfileDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type AccessTokenDTO struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func ToProfileDto(user models.User) ProfileDto {
	return ProfileDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
