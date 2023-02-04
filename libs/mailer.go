package libs

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(email, subject, link string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "PT. Rental Vehicle <example@gmail.com>")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", `<a href="`+link+`">Verify Email.</a>`)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(os.Getenv("SMPT_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
