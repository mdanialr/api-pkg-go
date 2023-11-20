package response

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// WithData option to add the given data to success response as `data` field.
func WithData(data any) AppOpt {
	return func(s *App) {
		s.Data = data
	}
}

// WithPaginate option to add the given paginate to success response as
// `pagination` field.
func WithPaginate(paginate any) AppOpt {
	return func(s *App) {
		s.Pagination = paginate
	}
}

// WithErr option to add given detail to error response as `error` field.
func WithErr(err error) AppOpt {
	return func(a *App) {
		var stdErr = new(Std)
		var valid validator.ValidationErrors

		switch {

		// error come from any layer that use std as error
		case errors.As(err, &stdErr):
			a.Code = stdErr.Code
			a.Message = stdErr.Message

		// error commonly come from validation in delivery layer
		case errors.As(err, &valid):
			a.Code = "ValidationError"
			a.Message = "Provided data is invalid, please check again"
			a.Error = NewValidationErrors(valid)

		// set default to unknown error
		default:
			a.Code = "UnknownError"
			a.Message = err.Error()
		}
	}
}
