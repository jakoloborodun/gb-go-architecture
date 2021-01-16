package smtpserv

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"

	"shop/models"
)

type SmtpServer interface {
	SendOrderConfirmationEmail(order *models.Order) error
	SendEmail(sendTo string, subject string, message []byte) error
}

type smtpServer struct {
	auth smtp.Auth
	host string
	port string
	from string
}

func (c *smtpServer) SendOrderConfirmationEmail(order *models.Order) error {
	subject := fmt.Sprintf("Your order #%d confirmed!", order.ID)
	message := fmt.Sprintf("Hi there! Thanks for your order #%d\nOur manager will contact you shortly to agree on a price at the phone number you specified: %s", order.ID, order.Phone)

	return c.SendEmail(order.Email, subject, []byte(message))
}

func (c *smtpServer) SendEmail(sendTo, subject string, message []byte) error {
	header := make(map[string]string)
	header["From"] = c.from
	header["To"] = sendTo
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	body := ""
	for key, value := range header {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	body += "\r\n" + base64.StdEncoding.EncodeToString(message)

	err := smtp.SendMail(c.host+":"+c.port, c.auth, c.from, []string{sendTo}, []byte(body))
	return err
}

// Create the authentication for the SendMail()
func NewSmtpAuth(identity, from, password, host string) smtp.Auth {
	log.Println(from)
	log.Println(password)
	log.Println(host)
	return smtp.PlainAuth(identity, from, password, host)
}

func NewSmtpServer(host, port, from, pass string) (SmtpServer, error) {
	auth := NewSmtpAuth("", from, pass, host)

	return &smtpServer{
		auth: auth,
		host: host,
		port: port,
		from: from,
	}, nil
}
