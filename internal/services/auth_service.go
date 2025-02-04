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

func LoginUser(user dtos.LoginDto) (dtos.AccessTokenDto, error) {
	userFromDB, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return dtos.AccessTokenDto{}, errors.New("not found")
	}

	if !auth.ComparePassword(user.Password, userFromDB.Password) {
		return dtos.AccessTokenDto{}, errors.New("unauthorized")
	}

	accessToken, expiresIn, tokenType, err := auth.GenerateJWT(userFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDto{}, errors.New("internal server error")
	}

	return dtos.AccessTokenDto{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		TokenType:   tokenType,
	}, nil
}
