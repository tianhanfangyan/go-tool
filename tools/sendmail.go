package tools

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/smtp"
	"os"
	"time"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func SendEmail(from, subject, msg, attach, host string, port int, username, password string, receviers ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", receviers...)
	m.SetHeader("Subject", subject)
	m.SetDateHeader("Date", time.Now())
	m.SetBody("text/plain", msg)

	if _, err := os.Stat(attach); err != nil {
		fmt.Println(err)
	}
	m.Attach(attach)

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	auth := unencryptedAuth{smtp.PlainAuth("", username, password, host)}
	d.Auth = auth

	return d.DialAndSend(m)
}
