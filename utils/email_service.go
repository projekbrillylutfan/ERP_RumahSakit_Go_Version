package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(mailDialer *gomail.Dialer, fromEmail string, toEmail string, name string, token string) error {
	verificationURL := fmt.Sprintf("http://localhost:3000/api/verify-email?token=%s", token)

	message := gomail.NewMessage()
	message.SetHeader("From", fromEmail)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", "Email Verification")
	message.SetBody("text/plain", fmt.Sprintf("Hi %s, please verify your email using this link: %s", name, verificationURL))

	// Send the email
	return mailDialer.DialAndSend(message)
}