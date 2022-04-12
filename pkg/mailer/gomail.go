package mailer

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type SMTPMailer struct {
	dialer *gomail.Dialer
}

func NewSMTPMailer(host string, port int, username, password string) *SMTPMailer {
	return &SMTPMailer{
		dialer: gomail.NewDialer(host, port, username, password),
	}
}

func (m SMTPMailer) SendEmail(from string, to []string, subject, body string) error {
	message := m.buildMessage(from, to, subject, body)
	return m.dialer.DialAndSend(message)
}

func (m SMTPMailer) buildMessage(from string, to []string, subject, body string) *gomail.Message {
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	return msg
}
func (m SMTPMailer) SetDevEnv() {
	m.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}
