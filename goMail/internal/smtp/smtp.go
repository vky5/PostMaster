package smtp

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
	"github.com/joho/godotenv"
)

func sendMails() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Access environment variables
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// Create a new message
	message := mail.NewMessage()

	// Define email headers
	var headers = map[string][]string{
		"From":    {"info@blacktree.in"},
		"To":      {"vky5@proton.me"},
		"Subject": {"Hello from Go Mail"},
	}

	// Set headers to the message
	message.SetHeaders(headers)

	// Set the email body
	message.SetBody("text/plain", "We are testing this but we are going to make it big")

	// Set up the SMTP dialer using the loaded environment variables
	dialer := mail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPass)

	// Send the email
	err = dialer.DialAndSend(message)
	if err != nil {
		// Print error message
		fmt.Printf("There was an error: %v\n", err)
	} else {
		// Print success message
		fmt.Println("Email sent successfully")
	}
}
