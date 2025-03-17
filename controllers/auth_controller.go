package controllers

import (
	"errors"
	"os"
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

	response := mapper.UserToLoginResponse(token)
	return &response, nil
}

func (c *AuthController) SignUp(request dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, error) {
	var existingUser models.User
	result := c.DB.Where("email = ?", request.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, errors.New("userExistsWithEmail")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("passwordProcessError")
	}

	newUser := mapper.RegisterRequestToUserModel(request, string(hashedPassword))

	// start DB Transaction
	tx := c.DB.Begin()

	err = tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("userRegistrationfailed")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("userRegistrationfailed")
	}

	response := mapper.UserToRegistrationResponse(newUser)

	return &response, nil
}

func (c *AuthController) ConfigureAdmin() error {
	firstName := os.Getenv("ADMIN_FIRST_NAME")
	lastName := os.Getenv("ADMIN_LAST_NAME")
	email := os.Getenv("ADMIN_EMAIL")
	password := os.Getenv("ADMIN_PASSWORD")

	var count int64

	c.DB.Model(&models.User{}).Where("email =?", email).Count(&count)
	if count > 0 {
		return errors.New("User exists!")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	admin := models.User{
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		Password:      string(hashedPassword),
		IsSuperuser:   true,
		Role:          models.Role("admin"),
		Status:        "active",
		EmailVerified: true,
	}

	err := c.DB.Create(&admin).Error

	return err
}
