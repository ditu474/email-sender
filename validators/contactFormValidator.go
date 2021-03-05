package validators

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/ditu474/email-sender/models"
)

//ContactFormValidator is the validator for the Contact Form
func ContactFormValidator(cf models.ContactForm) error {
	if govalidator.IsNull(cf.Email) {
		return fmt.Errorf("Required field: email")
	}
	if govalidator.IsNull(cf.Message) {
		return fmt.Errorf("Required field: message")
	}
	if govalidator.IsNull(cf.Name) {
		return fmt.Errorf("Required field: name")
	}
	if govalidator.IsNull(cf.Subject) {
		return fmt.Errorf("Required field: subject")
	}
	if !govalidator.IsEmail(cf.Email) {
		return fmt.Errorf("Invalid email")
	}
	return nil
}
