package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendMail(mail_to, body string) {
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASSWORD")

	fmt.Printf("%s: %s\n", from, pass)

	msg := "From: " + from + "\n" +
		"To: " + mail_to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail(fmt.Sprintf("%s:%s", os.Getenv("SMTP_SERVER"), os.Getenv("SMTP_PORT")),
		smtp.PlainAuth("", from, pass, os.Getenv("SMTP_SERVER")),
		from, []string{mail_to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
