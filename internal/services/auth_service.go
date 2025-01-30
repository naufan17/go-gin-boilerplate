package services

import (
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/models"
	"github.com/naufan17/go-gin-boilerplate/internal/repositories"
	"github.com/naufan17/go-gin-boilerplate/internal/utils"
	"golang.org/x/crypto/bcrypt"

	"errors"
)

func RegisterUser(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	if _, err := repositories.CreateUser(user); err != nil {
		return models.User{}, errors.New("conflict")
	}

	return user, nil
}

func LoginUser(user dtos.LoginDTO) (dtos.AccessTokenDTO, error) {
	userFromDB, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return dtos.AccessTokenDTO{}, errors.New("not found")
	}

	if !utils.ComparePassword(user.Password, userFromDB.Password) {
		return dtos.AccessTokenDTO{}, errors.New("unauthorized")
	}

	accessToken, expiresIn, tokenType, err := utils.GenerateJWT(userFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDTO{}, errors.New("internal server error")
	}

	return dtos.AccessTokenDTO{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		TokenType:   tokenType,
	}, nil
}
