package utils

import (
	"errors"
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

func ValidateJWT(tokenString string) (*config.Claims, error) {
	claims := &config.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexected signing method")
		}

		return config.JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Token expired")
	}

	return claims, nil
}
