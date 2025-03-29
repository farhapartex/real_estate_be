package dto

import "time"

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

type RegisterResponseDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
}

type UserDetailShortDTO struct {
	ID            uint       `json:"id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty"`
	EmailVerified bool       `json:"email_verified"`
	Role          string     `json:"role"`
	IsSuperuser   bool       `json:"is_superuser"`
	JoinedAt      time.Time  `json:"joined_at"`
	PhoneNumber   string     `json:"phone_number"`
	Website       string     `json:"website"`
	Status        string     `json:"status"`
}

type UserMeDTO struct {
	ID            uint       `json:"id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty"`
	EmailVerified bool       `json:"email_verified"`
	Role          string     `json:"role"`
}

type OwnerSignupRequestDTO struct {
	FirstName   string `json:"first_name" binding:"required,min=2,max=150"`
	LastName    string `json:"last_name" binding:"required,min=2,max=150"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8,max=72"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=20"`
}
