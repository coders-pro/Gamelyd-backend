package helper

import (
	"net/smtp"
	"fmt"

	templates "github.com/Gameware/templates"
)

func ForgotPasswordMail(receiver string, token string, username string) {
	
	// Sender data.
	from := "info@shipping-cargo.com"
	password := "ZzFV@#UqT7JAr"
  
	// Receiver email address.
	to := []string{
	  receiver,
	}
  
	// smtp server configuration.
	smtpHost := "shipping-cargo.com"
	smtpPort := "26"
  
	// Message.
	subject := "Subject: Reset Password\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := templates.ForgotPassword(token, username)
	msg := []byte(subject + mime + body)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println("Email Sent Successfully!")

  }