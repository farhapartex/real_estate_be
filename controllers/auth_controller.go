package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/farhapartex/real_estate_be/config"
	"github.com/farhapartex/real_estate_be/dto"
	"github.com/farhapartex/real_estate_be/lib/email"
	"github.com/farhapartex/real_estate_be/mapper"
	"github.com/farhapartex/real_estate_be/models"
	"github.com/farhapartex/real_estate_be/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB           *gorm.DB
	EmailService *email.EmailService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB:           db,
		EmailService: email.NewEmailService(),
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

func (c *AuthController) SignUp(request dto.OwnerSignupRequestDTO) (*dto.RegisterResponseDTO, error) {
	request.Email = strings.ToLower(request.Email)

	var existingUser models.User

	result := c.DB.Where("email = ?", request.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.EmailService.SendVerificationEmail(existingUser, "test")
		return nil, errors.New("userExistsWithEmail")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("passwordProcessError")
	}

	newUser := mapper.OwnerSignupDTOToUserModel(request, string(hashedPassword))

	// start DB Transaction
	tx := c.DB.Begin()

	err = tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("userRegistrationfailed")
	}

	ownerProfile := mapper.OwnerSignupDTOToProfileModel(request, newUser.ID)
	err = tx.Create(&ownerProfile).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("profileCreationFailed")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("userRegistrationfailed")
	}

	token, err := c.GenerateVerificationToken(newUser.ID)
	fmt.Println("token: ", token)

	c.EmailService.SendVerificationEmail(newUser, token)

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

func (c *AuthController) UserMeData(ctx *gin.Context) (*dto.UserMeDTO, error) {
	user, exists := ctx.Get(("user"))
	if !exists {
		return nil, errors.New("Authentication required")
	}
	userMode, ok := user.(models.User)
	if !ok {
		return nil, errors.New("Could not process user data")
	}

	response := mapper.UserToMeResponse(userMode)
	return &response, nil
}

func (c *AuthController) GenerateVerificationToken(userID uint) (string, error) {
	if c.DB == nil {
		c.DB = config.DB
	}

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	plainToken := base64.URLEncoding.EncodeToString(b)

	hash := sha256.Sum256([]byte(plainToken))
	hashedToken := hex.EncodeToString(hash[:])

	expiresAt := time.Now().Add(48 * time.Hour)

	verificationToken := models.VerificationToken{
		UserID:    userID,
		Token:     hashedToken,
		Type:      "email_verification",
		ExpiresAt: expiresAt,
	}

	if err := c.DB.Create(&verificationToken).Error; err != nil {
		return "", err
	}

	return plainToken, nil
}

func (c *AuthController) ResendVerification(email string) (bool, string, error) {
	// Find user by email
	var user models.User
	if err := c.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Don't reveal if the email exists or not for security
			return true, "If your account exists, a verification email will be sent", nil
		}
		return false, "Error processing request", err
	}

	if user.EmailVerified {
		return false, "Your account is already verified", nil
	}

	if err := c.DB.Where("user_id = ? AND type = ? AND used_at IS NULL AND expires_at > ?",
		user.ID, "email_verification", time.Now()).Delete(&models.VerificationToken{}).Error; err != nil {
		return false, "Error processing request", err
	}

	_, err := c.GenerateVerificationToken(user.ID)
	if err != nil {
		return false, "Failed to generate verification token", err
	}

	return true, "A new verification email has been sent to your address", nil
}

func (c *AuthController) VerifyAccount(token string) (bool, string, error) {
	// Hash the provided token to match against stored hash
	hash := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hash[:])

	// Find the token in the database
	var verificationToken models.VerificationToken
	if err := c.DB.Where("token = ? AND type = ?", hashedToken, "email_verification").First(&verificationToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, "Invalid verification token", nil
		}
		return false, "Error processing verification", err
	}

	if verificationToken.IsExpired() {
		return false, "Verification token has expired", nil
	}

	if verificationToken.IsUsed() {
		return false, "Verification token has already been used", nil
	}

	tx := c.DB.Begin()

	now := time.Now()
	if err := tx.Model(&models.User{}).Where("id = ?", verificationToken.UserID).Updates(map[string]interface{}{
		"email_verified": true,
		"verified_at":    now,
		"status":         "active",
	}).Error; err != nil {
		tx.Rollback()
		return false, "Failed to verify account", err
	}

	// Mark token as used
	verificationToken.MarkAsUsed()
	if err := tx.Save(&verificationToken).Error; err != nil {
		tx.Rollback()
		return false, "Failed to update verification token", err
	}

	if err := tx.Commit().Error; err != nil {
		return false, "Failed to complete verification", err
	}

	return true, "Your account has been successfully verified", nil
}
