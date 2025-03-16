package utils

import (
	"time"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &config.Claims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}
