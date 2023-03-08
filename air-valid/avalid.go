package avalid

import (
	"net/mail"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isvalidemail", isValidEmail)
	}
}

func isValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String();
	_, err := mail.ParseAddress(email)
  return err == nil
}