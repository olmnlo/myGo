package models

import (
	"fmt"
	"net/smtp"
)

func SendEmail(subject string, body Job, to string) error {
	// Set up the email details
	from := "tt6354330@gmail.com"
	password := "cidclyodkwoqfrde" // Use an app password for Gmail if you have 2FA enabled
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Create the authentication for the email
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Compose the email message
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		"msg:" + body.Msg + "\n" +
		"TimeEnd: " + body.TimeEnd + "\n" +
		"TimeStart: " + body.TimeStart + "\n" +
		"TimeDuration: " + body.TimeDuration)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
