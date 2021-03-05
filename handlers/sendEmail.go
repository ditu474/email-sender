package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"

	"github.com/ditu474/email-sender/models"
	"github.com/ditu474/email-sender/validators"
)

// SendEmail is the handler responsible for sending the email
type SendEmail struct {
	l *log.Logger
}

// NewSendEmail is the function responsible of creating a new SendEmail struct
func NewSendEmail(l *log.Logger) *SendEmail {
	return &SendEmail{
		l: l,
	}
}

func (s *SendEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		s.postHandler(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *SendEmail) postHandler(rw http.ResponseWriter, r *http.Request) {
	cf := models.ContactForm{}
	err := json.NewDecoder(r.Body).Decode(&cf)
	if err != nil && err != io.EOF {
		s.l.Printf("Error decoding request: %v\n", err)

		// WriteHeader must to go before Write
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"error":"Bad JSON"}`))
		return
	}

	if err := validators.ContactFormValidator(cf); err != nil {
		s.l.Printf("Validation fails: %v\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, err)))
		return
	}
	err = s.sendEmail(rw, cf)
	if err != nil {
		s.l.Printf("Error sending email: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"The mail could not be sent"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message":"Email sent"}`))
}

func (s *SendEmail) sendEmail(rw http.ResponseWriter, cf models.ContactForm) error {
	ha := &hostAuth{
		"smtp.mailtrap.io",
		"25",
		"",
		"",
	}
	auth := smtp.PlainAuth("", ha.username, ha.password, ha.smtpHost)
	from := "danielrg0322@gmail.com"
	to := []string{cf.Email}
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"Hello %s,\r\n%s\r\n", cf.Email, cf.Subject, cf.Name, cf.Message))
	err := smtp.SendMail(ha.smtpHost+":"+ha.smtPort, auth, from, to, msg)
	return err
}

type hostAuth struct {
	smtpHost string
	smtPort  string
	username string
	password string
}
