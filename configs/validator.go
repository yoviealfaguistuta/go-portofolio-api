package configs

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// RegisterTagNameFunc registers a function to get alternate names for StructFields.
	// eg. Title become title, CreatedAt become created_at
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Custom validation for uuid.UUID fields.
	/*_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})*/

	return validate
}
