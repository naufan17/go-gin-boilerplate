package services

import (
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
	"github.com/naufan17/go-gin-boilerplate/internal/repositories"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"

	"errors"
)

func RegisterUser(user dtos.RegisterDto) (models.User, error) {
	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		return models.User{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	userSaveToDB := dtos.ToUserModel(user)

	if _, err := repositories.CreateUser(userSaveToDB); err != nil {
		return models.User{}, errors.New("conflict")
	}

	return userSaveToDB, nil
}

func LoginUser(user dtos.LoginDto, ipAddress string, userAgent string) (dtos.AccessTokenDto, dtos.RefreshTokenDto, error) {
	userFromDB, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return dtos.AccessTokenDto{}, dtos.RefreshTokenDto{}, errors.New("not found")
	}

	if !auth.ComparePassword(user.Password, userFromDB.Password) {
		return dtos.AccessTokenDto{}, dtos.RefreshTokenDto{}, errors.New("unauthorized")
	}

	accessAccessToken, accessExpiresIn, accessTokenType, err := auth.GenerateJWTAccess(userFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDto{}, dtos.RefreshTokenDto{}, errors.New("internal server error")
	}

	sessionFromDB, err := repositories.CreateSession(userFromDB.ID, ipAddress, userAgent)

	if err != nil {
		return dtos.AccessTokenDto{}, dtos.RefreshTokenDto{}, errors.New("internal server error")
	}

	refreshRefreshToken, refreshExpiresIn, refreshtTokenType, err := auth.GenerateJWTRefresh(sessionFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDto{}, dtos.RefreshTokenDto{}, errors.New("internal server error")
	}

	return dtos.AccessTokenDto{
			AccessToken: accessAccessToken,
			ExpiresIn:   accessExpiresIn,
			TokenType:   accessTokenType,
		},
		dtos.RefreshTokenDto{
			RefreshToken: refreshRefreshToken,
			ExpiresIn:    refreshExpiresIn,
			TokenType:    refreshtTokenType,
		}, nil
}
