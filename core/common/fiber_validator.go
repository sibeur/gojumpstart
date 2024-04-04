package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type FiberValidator struct {
	validator *validator.Validate
}

func NewFiberValidator() *FiberValidator {
	return &FiberValidator{
		validator: validator.New(),
	}
}

func (v *FiberValidator) Validate(data any) []FiberErrorMessage {
	validationErrors := []FiberErrorMessage{}
	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem FiberErrorMessage

			elem.Key = err.Field()
			elem.Value = getErrorMessage(err)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is a required field", e.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", e.Field())
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s", e.Field(), e.Param())
	case "nefield":
		return fmt.Sprintf("%s must not be equal to %s", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", e.Field(), e.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", e.Field(), e.Param())
	case "eq":
		return fmt.Sprintf("%s must be equal to %s", e.Field(), e.Param())
	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("%s must be of length %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must have a minimum value of %s", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must have a maximum value of %s", e.Field(), e.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", e.Field())
	case "uri":
		return fmt.Sprintf("%s must be a valid URI", e.Field())
	case "alpha":
		return fmt.Sprintf("%s must contain only letters", e.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", e.Field())
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", e.Field())
	case "hexadecimal":
		return fmt.Sprintf("%s must be a valid hexadecimal number", e.Field())
	case "hexcolor":
		return fmt.Sprintf("%s must be a valid hex color", e.Field())
	case "rgb":
		return fmt.Sprintf("%s must be a valid RGB color", e.Field())
	case "rgba":
		return fmt.Sprintf("%s must be a valid RGBA color", e.Field())
	case "hsl":
		return fmt.Sprintf("%s must be a valid HSL color", e.Field())
	case "hsla":
		return fmt.Sprintf("%s must be a valid HSLA color", e.Field())
	case "json":
		return fmt.Sprintf("%s must be a valid JSON string", e.Field())
	case "base64":
		return fmt.Sprintf("%s must be a valid Base64 string", e.Field())
	case "ip":
		return fmt.Sprintf("%s must be a valid IP address", e.Field())
	case "ipv4":
		return fmt.Sprintf("%s must be a valid IPv4 address", e.Field())
	case "ipv6":
		return fmt.Sprintf("%s must be a valid IPv6 address", e.Field())
	case "ssn":
		return fmt.Sprintf("%s must be a valid SSN", e.Field())
	case "isbn":
		return fmt.Sprintf("%s must be a valid ISBN", e.Field())
	case "isbn10":
		return fmt.Sprintf("%s must be a valid ISBN-10", e.Field())
	case "isbn13":
		return fmt.Sprintf("%s must be a valid ISBN-13", e.Field())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", e.Field())
	case "uuid3":
		return fmt.Sprintf("%s must be a valid version 3 UUID", e.Field())
	case "uuid4":
		return fmt.Sprintf("%s must be a valid version 4 UUID", e.Field())
	case "uuid5":
		return fmt.Sprintf("%s must be a valid version 5 UUID", e.Field())
	case "ascii":
		return fmt.Sprintf("%s must contain only ASCII characters", e.Field())
	case "printascii":
		return fmt.Sprintf("%s must contain only printable ASCII characters", e.Field())
	case "multibyte":
		return fmt.Sprintf("%s must contain one or more multibyte characters", e.Field())
	case "datauri":
		return fmt.Sprintf("%s must be a valid Data URI", e.Field())
	case "latitude":
		return fmt.Sprintf("%s must be a valid latitude", e.Field())
	case "longitude":
		return fmt.Sprintf("%s must be a valid longitude", e.Field())
	case "number":
		return fmt.Sprintf("%s must be a valid number", e.Field())
	case "lowercase":
		return fmt.Sprintf("%s must be in lowercase", e.Field())
	case "uppercase":
		return fmt.Sprintf("%s must be in uppercase", e.Field())
	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime", e.Field())
	case "creditcard":
		return fmt.Sprintf("%s must be a valid credit card number", e.Field())
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", e.Field(), e.Param())
	case "unique":
		return fmt.Sprintf("%s must be unique", e.Field())
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}
