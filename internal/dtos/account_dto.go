package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
)

type ProfileDto struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicture string    `json:"profile_picture"`
	IsVerified     bool      `json:"is_verified"`
}

type SessionDto struct {
	ID           uuid.UUID `json:"id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	Status       string    `json:"status"`
	LoginAt      time.Time `json:"login_at"`
	LastActiveAt time.Time `json:"last_active_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type UpdateProfileDto struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type UpdatePasswordDto struct {
	Password        string `json:"password" validate:"required,min=10"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func ToProfileDto(user models.User) ProfileDto {
	return ProfileDto{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		ProfilePicture: user.ProfilePicture,
		IsVerified:     user.IsVerified,
	}
}

func ToSessionDto(session models.Session) SessionDto {
	return SessionDto{
		ID:           session.ID,
		IPAddress:    session.IPAddress,
		UserAgent:    session.UserAgent,
		Status:       getStatus(session.ExpiresAt),
		LoginAt:      session.LoginAt,
		LastActiveAt: session.LastActiveAt,
		ExpiresAt:    session.ExpiresAt,
	}
}

func getStatus(expiresAt time.Time) string {
	if time.Now().After(expiresAt) {
		return "expired"
	}

	return "active"
}
