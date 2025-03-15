package config

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
