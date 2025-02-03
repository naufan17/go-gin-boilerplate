package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
)

var (
	jwtSecret = []byte(configs.LoadConfig().JWTSecret)
	jwtExpStr = configs.LoadConfig().JWTExp
	jwtExp, _ = strconv.Atoi(jwtExpStr)
)

type Claims struct {
	Sub uuid.UUID `json:"sub"`
	Iat int64     `json:"iat"`
	jwt.StandardClaims
}

func GenerateJWT(id uuid.UUID) (string, int64, string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtExp) * time.Millisecond)

	claims := &Claims{
		Sub: id,
		Iat: time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", 0, "", err
	}

	return tokenString, expirationTime.Unix(), "Bearer", nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
