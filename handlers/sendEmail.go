package handlers

import (
	"net/http"
)

// SendEmail is the handler responsible for sending the email
type SendEmail struct{}

// NewSendEmail is the function responsible of creating a new SendEmail struct
func NewSendEmail() *SendEmail {
	return &SendEmail{}
}

func (s *SendEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.sendEmail(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *SendEmail) sendEmail(rw http.ResponseWriter, r *http.Request) {

}
