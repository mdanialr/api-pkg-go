package response

import (
	"sync"

	help "github.com/mdanialr/api-pkg-go/helper"

	"github.com/go-playground/validator/v10"
)

var once sync.Once
var bagOfErrorMessages map[string]func(fe validator.FieldError) string

// AddErrorMessage add or replace error message mapping to bag from given key
// value.
func AddErrorMessage(k string, v func(validator.FieldError) string) {
	once.Do(func() {
		bagOfErrorMessages = make(map[string]func(fe validator.FieldError) string)
		fillUpDefaultBagOfErrMessages()
	})
	bagOfErrorMessages[k] = v
}

// fillUpDefaultBagOfErrMessages fill in bag of error messages with default
// message mapping that handle common usage.
func fillUpDefaultBagOfErrMessages() {
	bagOfErrorMessages["required"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should not be empty"
	}
	bagOfErrorMessages["len"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have length " + fe.Param() + " characters"
	}
	bagOfErrorMessages["numeric"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should only contain numeric characters"
	}
	bagOfErrorMessages["max"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have value or length characters same or less than " + fe.Param()
	}
	bagOfErrorMessages["min"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have value or length characters same or more than " + fe.Param()
	}
	bagOfErrorMessages["url"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should be valid url and have FQDN"
	}
	bagOfErrorMessages["boolean"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should only contain either (true | false)"
	}
	bagOfErrorMessages["alpha"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should only contain alphabet characters"
	}
	bagOfErrorMessages["alphanum"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should only contain alphabet and or numeric characters"
	}
	bagOfErrorMessages["image"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should only contain one of (jpg|jpeg|png) file extension"
	}
	bagOfErrorMessages["gte"] = func(fe validator.FieldError) string {
		// gte: is comparison only for number
		return "Param " + help.ToSnake(fe.Field()) + " should have value same or more than " + fe.Param()
	}
	bagOfErrorMessages["ltefield"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have value that's more than the value of param " + help.ToSnake(fe.Param())
	}
	bagOfErrorMessages["ltfield"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have value that's equal or more than the value of param " + help.ToSnake(fe.Param())
	}
	bagOfErrorMessages["gtfield"] = func(fe validator.FieldError) string {
		return "Param " + help.ToSnake(fe.Field()) + " should have value that's equal or less than the value of param " + help.ToSnake(fe.Param())
	}
}

// validationError standard object that hold error from validation.
type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewValidationErrors return ready to display error validations.
func NewValidationErrors(validations validator.ValidationErrors) []validationError {
	once.Do(func() {
		bagOfErrorMessages = make(map[string]func(fe validator.FieldError) string)
		fillUpDefaultBagOfErrMessages()
	})
	var validErr []validationError
	for _, valid := range validations {
		validErr = append(validErr, validationError{
			Field:   help.ToSnake(valid.Field()),
			Message: errMsgMapping(valid),
		})

	}
	return validErr
}

// errMsgMapping custom error message constructor from validator.
func errMsgMapping(fe validator.FieldError) string {
	if _, ok := bagOfErrorMessages[fe.Tag()]; ok {
		return bagOfErrorMessages[fe.Tag()](fe)
	}
	return fe.Error()
}
