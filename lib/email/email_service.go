package email

import (
	"fmt"

	"github.com/farhapartex/real_estate_be/models"
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) SendVerificationEmail(user models.User) error {
	// This function will be implemented later
	// For now, just log that we would send an email
	fmt.Printf("Would send verification email to: %s (%s)\n", user.Email, user.FirstName+" "+user.LastName)
	return nil
}

// GenerateVerificationToken creates a token for email verification
// This is a placeholder function that will be implemented later
func (es *EmailService) GenerateVerificationToken(userID uint) (string, error) {
	// This will be implemented later
	// For now, return a placeholder
	return fmt.Sprintf("placeholder-token-%d", userID), nil
}
