package libs

import (
	"os"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL      string
	Username string
	Subject  string
}

func SendMail(user *models.User, data *EmailData) error {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "PT. Rental Vehicle",
			Link: os.Getenv("BASE_URL"),
			Logo: "https://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	// Create an email body using hermes
	emailBody, err := h.GenerateHTML(hermes.Email{
		Body: hermes.Body{
			Name: user.Username,
			Intros: []string{
				"Welcome to PT. Rental Vehicle!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please click the following button to verify your email. This link expires in 30 minutes.",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  data.URL,
					},
				},
			},
		},
	})

	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "PT. Rental Vehicle <example@gmail.com>")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", emailBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(os.Getenv("SMPT_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
