package calendoo

import (
	"net/smtp"
)

type Email struct {
	To      []string
	From    string
	Subject string
	Body    string
}

type Mailer struct {
	Auth smtp.Auth
	Host string
	Port string
}

func (m *Mailer) SendMail(e *Email) error {
	header := make(map[string]string)
	header["From"] = e.From
	header["To"] = e.To[0]
	header["Subject"] = e.Subject

	message := ""
	for k, v := range header {
		message += k + ": " + v + "\r\n"
	}
	message += "\r\n" + e.Body

	err := smtp.SendMail(m.Host+":"+m.Port, m.Auth, e.From, e.To, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
