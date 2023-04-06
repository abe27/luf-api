package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	mail "gopkg.in/mail.v2"
)

func SendMail(mail_to, subject, body string) {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", mail_to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := mail.NewDialer(os.Getenv("SMTP_SERVER"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
