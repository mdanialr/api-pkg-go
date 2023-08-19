package response

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// WithData option to add the given data to success response as `data` field.
func WithData(data any) SuccessOption {
	return func(s *AppSuccess) {
		s.Data = data
	}
}

// WithPaginate option to add the given paginate to success response as
// `pagination` field.
func WithPaginate(paginate any) SuccessOption {
	return func(s *AppSuccess) {
		s.Pagination = paginate
	}
}

// WithErr option to add given detail to error response as `error` field.
func WithErr(err error) AppErrorOption {
	return func(a *AppError) {
		// error come from use case layer that use std error
		var stdErr = new(Std)
		if errors.As(err, &stdErr) {
			a.Code = stdErr.Code
			a.Message = stdErr.Message
			return
		}
		// error come from validation in delivery layer
		var valid validator.ValidationErrors
		if errors.As(err, &valid) {
			a.Code = "ValidationError"
			a.Message = "Provided data is invalid, please check again"
			a.Error = NewValidationErrors(valid)
			return
		}
		// set default to unknown error
		a.Code = "UnknownError"
		a.Message = err.Error()
	}
}
