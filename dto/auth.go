package dto

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string             `json:"token"`
	User  UserDetailShortDTO `json:"user"`
}

type RegisterRequestDTO struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type RegisterResponseDTO struct {
	FirstName string `json:"first_name"`
}

type UserDetailShortDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
