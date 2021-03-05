package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ditu474/email-sender/models"
	"github.com/ditu474/email-sender/validators"
)

// SendEmail is the handler responsible for sending the email
type SendEmail struct {
	l         *log.Logger
	validator validators.Validator
}

// NewSendEmail is the function responsible of creating a new SendEmail struct
func NewSendEmail(l *log.Logger, validator validators.Validator) *SendEmail {
	return &SendEmail{
		l:         l,
		validator: validator,
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
	defer r.Body.Close()

	//Save the request body
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	//Because we read the io.ReadCloser, the content has gone
	//so we set it back
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	cf := models.ContactForm{}
	err := json.NewDecoder(r.Body).Decode(&cf)
	if err != nil && err != io.EOF {
		s.l.Printf("Error decoding request: %v\n", err)

		// WriteHeader must to go before Write
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"error":"Bad JSON"}`))
		return
	}

	//Because we read the body again, we set it back for the validator
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := s.validator.Validation(r); err != nil {
		s.l.Printf("Validation fails: %v\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, err)))
		return
	}
	s.sendEmail(rw, cf)
}

func (s *SendEmail) sendEmail(rw http.ResponseWriter, cf models.ContactForm) {
	fmt.Fprint(rw, cf)
}
