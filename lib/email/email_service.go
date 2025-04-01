package email

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/farhapartex/real_estate_be/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailType string

const (
	VerificationEmail  EmailType = "verification"
	ResetPasswordEmail EmailType = "reset_password"
	WelcomeEmail       EmailType = "welcome"
	NotificationEmail  EmailType = "notification"
)

type EmailData struct {
	RecipientName    string
	RecipientEmail   string
	Subject          string
	VerificationLink string
	ResetLink        string
	CompanyName      string
	SupportEmail     string
	ExpiryTime       string
	CustomData       map[string]interface{} // For any additional data
}

type EmailService struct {
	client        *sendgrid.Client
	senderEmail   string
	senderName    string
	templatePath  string
	frontendURL   string
	supportEmail  string
	companyName   string
	tokenValidity time.Duration
}

func NewEmailService() *EmailService {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderName := os.Getenv("SENDER_NAME")
	templatePath := os.Getenv("EMAIL_TEMPLATE_PATH")
	frontendURL := os.Getenv("FRONTEND_URL")
	supportEmail := os.Getenv("SUPPORT_EMAIL")
	companyName := os.Getenv("COMPANY_NAME")

	// Default values if environment variables are not set
	if senderEmail == "" {
		senderEmail = "noreply@yourdomain.com"
	}
	if senderName == "" {
		senderName = "Your Real Estate App"
	}
	if templatePath == "" {
		templatePath = "templates/emails"
	}
	if frontendURL == "" {
		frontendURL = "https://ghor.online"
	}
	if supportEmail == "" {
		supportEmail = "support@yourdomain.com"
	}
	if companyName == "" {
		companyName = "Your Real Estate Company"
	}

	return &EmailService{
		client:        sendgrid.NewSendClient(apiKey),
		senderEmail:   senderEmail,
		senderName:    senderName,
		templatePath:  templatePath,
		frontendURL:   frontendURL,
		supportEmail:  supportEmail,
		companyName:   companyName,
		tokenValidity: 24 * time.Hour, // Default token validity: 24 hours
	}
}

func (es *EmailService) SendEmail(recipientEmail, recipientName, subject, htmlContent, plainContent string) error {
	from := mail.NewEmail(es.senderName, es.senderEmail)
	to := mail.NewEmail(recipientName, recipientEmail)

	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)
	response, err := es.client.Send(message)

	if err != nil {
		return errors.New("Error to send email: " + err.Error())
	}

	if response.StatusCode >= 400 {
		return errors.New("sendgrid API error: status code" + string(response.StatusCode))
	}

	return nil
}

func (es *EmailService) RenderTemplate(templateName string, data EmailData) (string, string, error) {
	// Path to HTML template
	htmlPath := filepath.Join(es.templatePath, templateName, "html.tmpl")
	// Path to plaintext template
	textPath := filepath.Join(es.templatePath, templateName, "text.tmpl")

	// Parse HTML template
	htmlTmpl, err := template.ParseFiles(htmlPath)
	if err != nil {
		return "", "", fmt.Errorf("error parsing HTML template: %w", err)
	}

	// Parse text template
	textTmpl, err := template.ParseFiles(textPath)
	if err != nil {
		return "", "", fmt.Errorf("error parsing text template: %w", err)
	}

	// Render HTML template
	var htmlBuffer bytes.Buffer
	if err := htmlTmpl.Execute(&htmlBuffer, data); err != nil {
		return "", "", fmt.Errorf("error executing HTML template: %w", err)
	}

	// Render text template
	var textBuffer bytes.Buffer
	if err := textTmpl.Execute(&textBuffer, data); err != nil {
		return "", "", fmt.Errorf("error executing text template: %w", err)
	}

	return htmlBuffer.String(), textBuffer.String(), nil
}

func (es *EmailService) SendVerificationEmail(user models.User, token string) error {
	// token, err := es.GenerateVerificationToken(user.ID)
	// if err != nil {
	// 	return fmt.Errorf("error generating verification token: %w", err)
	// }
	fmt.Println("working to send email")
	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", es.frontendURL, token)
	expiryTime := time.Now().Add(es.tokenValidity).Format("Jan 2, 2006 at 3:04 PM")

	data := EmailData{
		RecipientName:    user.FirstName + " " + user.LastName,
		RecipientEmail:   user.Email,
		Subject:          "Verify Your Email Address",
		VerificationLink: verificationLink,
		CompanyName:      es.companyName,
		SupportEmail:     es.supportEmail,
		ExpiryTime:       expiryTime,
	}

	htmlContent, plainContent, err := es.RenderTemplate(string(VerificationEmail), data)
	if err != nil {
		return fmt.Errorf("error rendering verification email template: %w", err)
	}

	err = es.SendEmail(user.Email, data.RecipientName, data.Subject, htmlContent, plainContent)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	fmt.Println("Email sent!!")

	return nil
}

// SendPasswordResetEmail sends password reset email
func (es *EmailService) SendPasswordResetEmail(user models.User) error {
	token, err := es.GenerateResetToken(user.ID)
	if err != nil {
		return fmt.Errorf("error generating reset token: %w", err)
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", es.frontendURL, token)
	expiryTime := time.Now().Add(es.tokenValidity).Format("Jan 2, 2006 at 3:04 PM")

	data := EmailData{
		RecipientName:  user.FirstName + " " + user.LastName,
		RecipientEmail: user.Email,
		Subject:        "Reset Your Password",
		ResetLink:      resetLink,
		CompanyName:    es.companyName,
		SupportEmail:   es.supportEmail,
		ExpiryTime:     expiryTime,
	}

	htmlContent, plainContent, err := es.RenderTemplate(string(ResetPasswordEmail), data)
	if err != nil {
		return fmt.Errorf("error rendering password reset email template: %w", err)
	}

	return es.SendEmail(user.Email, data.RecipientName, data.Subject, htmlContent, plainContent)
}

// SendWelcomeEmail sends a welcome email to newly verified users
func (es *EmailService) SendWelcomeEmail(user models.User) error {
	data := EmailData{
		RecipientName:  user.FirstName + " " + user.LastName,
		RecipientEmail: user.Email,
		Subject:        "Welcome to " + es.companyName,
		CompanyName:    es.companyName,
		SupportEmail:   es.supportEmail,
	}

	htmlContent, plainContent, err := es.RenderTemplate(string(WelcomeEmail), data)
	if err != nil {
		return fmt.Errorf("error rendering welcome email template: %w", err)
	}

	return es.SendEmail(user.Email, data.RecipientName, data.Subject, htmlContent, plainContent)
}

// SendCustomEmail sends a custom email based on provided template and data
func (es *EmailService) SendCustomEmail(recipientEmail, recipientName, subject string,
	templateName string, customData map[string]interface{}) error {

	data := EmailData{
		RecipientName:  recipientName,
		RecipientEmail: recipientEmail,
		Subject:        subject,
		CompanyName:    es.companyName,
		SupportEmail:   es.supportEmail,
		CustomData:     customData,
	}

	htmlContent, plainContent, err := es.RenderTemplate(templateName, data)
	if err != nil {
		return fmt.Errorf("error rendering custom email template: %w", err)
	}

	return es.SendEmail(recipientEmail, recipientName, subject, htmlContent, plainContent)
}

// GenerateVerificationToken creates a token for email verification
func (es *EmailService) GenerateVerificationToken(userID uint) (string, error) {
	// Generate a random token
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(b)

	// Store token in database associated with the user and set expiry
	// This is a placeholder - implement your database storage logic here
	// Example: es.tokenRepository.StoreToken(token, userID, "verification", time.Now().Add(es.tokenValidity))

	return token, nil
}

// GenerateResetToken creates a token for password reset
func (es *EmailService) GenerateResetToken(userID uint) (string, error) {
	// Generate a random token
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(b)

	// Store token in database associated with the user and set expiry
	// This is a placeholder - implement your database storage logic here
	// Example: es.tokenRepository.StoreToken(token, userID, "reset", time.Now().Add(es.tokenValidity))

	return token, nil
}

// VerifyToken validates a token and returns the associated user ID
func (es *EmailService) VerifyToken(token, tokenType string) (uint, error) {
	// Placeholder - implement your database lookup logic here
	// Example: return es.tokenRepository.VerifyToken(token, tokenType)
	return 0, fmt.Errorf("not implemented")
}

// SetTokenValidity allows customizing token expiration time
func (es *EmailService) SetTokenValidity(duration time.Duration) {
	es.tokenValidity = duration
}
