package controllers

import (
	"errors"
	"time"

	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
	"github.com/farhapartex/real_estate_be/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

func (c *AuthController) Login(request dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	var user models.User
	result := c.DB.Where("email = ?", request.Email).First(&user)

	if result.Error != nil {
		return nil, errors.New("invalidCredentials")
	}

	if user.Status != "active" {
		return nil, errors.New("accountNotActive")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalidCredentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("unableGenerateToken")
	}

	currentTime := time.Now()
	c.DB.Model(&user).Update("last_login_at", currentTime)

	response := mapper.UserToLoginResponse(user, token)
	return &response, nil
}
