package email

import (
	"email_sender/config"

	"gopkg.in/gomail.v2"
)

func SendEmail(from, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.yandex.ru", 465, config.SMTPUser, config.SMTPPass)
	d.SSL = true

	return d.DialAndSend(m)
}
