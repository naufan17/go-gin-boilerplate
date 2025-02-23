package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/naufan17/go-gin-boilerplate/config"
)

var (
	jwtAccessSecret  = []byte(config.LoadConfig().JWTAccessSecret)
	jwtRefreshSecret = []byte(config.LoadConfig().JWTRefreshSecret)
	jwtAccessExpStr  = config.LoadConfig().JWTAccessExp
	jwtRefreshExpStr = config.LoadConfig().JWTRefreshExp
	jwtAccessExp, _  = strconv.Atoi(jwtAccessExpStr)
	jwtRefreshExp, _ = strconv.Atoi(jwtRefreshExpStr)
)

type Claims struct {
	Sub uuid.UUID `json:"sub"`
	Iat int64     `json:"iat"`
	jwt.StandardClaims
}

func GenerateJWTAccess(id uuid.UUID) (string, int64, string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtAccessExp) * time.Millisecond)

	claims := &Claims{
		Sub: id,
		Iat: time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtAccessSecret)

	if err != nil {
		return "", 0, "", err
	}

	return tokenString, expirationTime.Unix(), "Bearer", nil
}

func GenerateJWTRefresh(id uuid.UUID) (string, int64, string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtRefreshExp) * time.Millisecond)

	claims := &Claims{
		Sub: id,
		Iat: time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtRefreshSecret)

	if err != nil {
		return "", 0, "", err
	}

	return tokenString, expirationTime.Unix(), "Bearer", nil
}

func ValidateJWTAccess(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtAccessSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func ValidateJWTRefresh(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtRefreshSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
