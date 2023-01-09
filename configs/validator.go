package configs

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate
}
