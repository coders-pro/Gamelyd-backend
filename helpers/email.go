package helper

import (
	"fmt"

	templates "github.com/Gameware/templates"
	mail "github.com/xhit/go-simple-mail/v2"
)

func ForgotPasswordMail(receiver string, token string, username string) {

	server := mail.NewSMTPClient()
	server.Host = "gamelyd.co"
	server.Port = 587
	server.Username = "no-reply@gamelyd.co"
	server.Password = "gamelydpass"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("Mcbobby from Gamelyd <no-reply@gamelyd.co>")
	email.AddTo(receiver)
	email.SetSubject("Reset Password")

	email.SetBody(mail.TextHTML, templates.ForgotPassword(token, username))
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println(err)
	}

	
	// Sender data.
	// from := "no-reply@gamelyd.co"
	// password := "gamelydpass"
  
	// // Receiver email address.
	// to := []string{
	//   receiver,
	// }
  
	// smtp server configuration.
	// smtpHost := "gamelyd.co"
	// smtpPort := "26"
  
	// // Message.
	// subject := "Subject: Reset Password\n"
	// mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// body := templates.ForgotPassword(token, username)
	// msg := []byte(subject + mime + body)
	// // Authentication.
	// auth := smtp.PlainAuth("", from, password, smtpHost)
	
	// // Sending email.
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	// if err != nil {
	//   fmt.Println(err)
	//   return
	// }
	// fmt.Println("Email Sent Successfully!")

  }

  func SendEmail(receiver string, template string, subject string) {


	server := mail.NewSMTPClient()
	server.Host = "gamelyd.co"
	server.Port = 587
	server.Username = "no-reply@gamelyd.co"
	server.Password = "gamelydpass"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("Mcbobby from Gamelyd <no-reply@gamelyd.co>")
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