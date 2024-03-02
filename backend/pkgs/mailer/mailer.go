// Package mailer provides a simple interface to send emails using SMTP.
package mailer

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"strconv"
)

type Mailer struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	From     string `json:"from,omitempty"`
}

func (m *Mailer) Ready() bool {
	return m.Host != "" && m.Port != 0 && m.Username != "" && m.Password != "" && m.From != ""
}

func (m *Mailer) server() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}

func (m *Mailer) Send(msg Message) error {
	server := m.server()

	header := map[string]string{
		"From":                      m.From,
		"Subject":                   mime.QEncoding.Encode("UTF-8", msg.Subject),
		"MIME-Version":              "1.0",
		"Content-Type":              "text/html; charset=\"utf-8\"",
		"Content-Transfer-Encoding": "base64",
	}

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(msg.Body))

	return smtp.SendMail(
		server,
		smtp.PlainAuth("", m.Username, m.Password, m.Host),
		m.From,
		msg.ToAddresses(),
		[]byte(message),
	)
}
