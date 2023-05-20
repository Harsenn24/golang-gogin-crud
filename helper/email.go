package helper

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) (string, error) {
    // Set up email configuration
    email := gomail.NewMessage()
    email.SetHeader("From", os.Getenv("SENDER_EMAIL"))
    email.SetHeader("To", to)
    email.SetHeader("Subject", subject)
    email.SetBody("text/plain", body)

    port_email, err := strconv.Atoi( os.Getenv("PORT_EMAIL"))
    if err != nil {
		return "Error parsing port", nil
	}

    // Set up SMTP server details
    d := gomail.NewDialer(os.Getenv("MAIL_TRANSPORT"), port_email, os.Getenv("USER_EMAIL"), os.Getenv("PASS_EMAIL"))

    // Authenticate with Gmail account
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Use this line if you encounter certificate issues

    // Send the email
    if err := d.DialAndSend(email); err != nil {
        return "failed send email", nil
    }

    return "Ok", nil
}

