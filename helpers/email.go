package helper

import (
	"fmt"
	"net/smtp"

	templates "github.com/Gameware/templates"
	mail "github.com/xhit/go-simple-mail/v2"
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

  func SendEmail(receiver string, template string, subject string) {


	server := mail.NewSMTPClient()
	server.Host = "shipping-cargo.com"
	server.Port = 587
	server.Username = "info@shipping-cargo.com"
	server.Password = "ZzFV@#UqT7JAr"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("Mcbobby from Gamelyd <info@shipping-cargo.com>")
	email.AddTo(receiver)
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, template)
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println(err)
	}


  }