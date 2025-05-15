package config

import "os"

var (
	FromEmail string
	SMTPUser  string
	SMTPPass  string
)

func Init() {
	FromEmail = os.Getenv("SMTP_FROM")
	SMTPUser = os.Getenv("SMTP_USER")
	SMTPPass = os.Getenv("SMTP_PASS")
}
