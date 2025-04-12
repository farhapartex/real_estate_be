package email

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(to, subject, body string) error {

	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY environment variable not set")
	}

	from := os.Getenv("SENDER_EMAIL")
	if from == "" {
		from = "noreply@yourdomain.com" // Default sender
	}
	senderName := os.Getenv("SENDER_NAME")
	if senderName == "" {
		senderName = "Your Application" // Default sender name
	}

	message := mail.NewSingleEmail(
		mail.NewEmail(senderName, from),
		subject,
		mail.NewEmail("", to),
		body, // Plain text version
		body, // HTML version (same as plain text in this case)
	)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("error sending email, status code: %d, body: %s",
			response.StatusCode, response.Body)
	}

	log.Printf("Email sent successfully to %s, status code: %d", to, response.StatusCode)
	return nil
}
