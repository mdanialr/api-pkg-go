package response

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestWithData(t *testing.T) {
	var obj struct {
		Username string
	}
	obj.Username = "user"

	testCases := []struct {
		name       string
		sampleData any
		expect     func(data any) AppSuccess
	}{
		{
			name: "Given calling WithData and give it data object {'username':'user'} Should return that" +
				" as data object",
			sampleData: obj,
			expect: func(data any) AppSuccess {
				return AppSuccess{Data: data}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(AppSuccess)
			WithData(tc.sampleData)(app)
			assert.Equal(t, tc.expect(tc.sampleData), *app)
		})
	}
}

func TestWithPaginate(t *testing.T) {
	var obj struct {
		Limit int
	}
	obj.Limit = 10

	testCases := []struct {
		name       string
		sampleData any
		expect     func(pg any) AppSuccess
	}{
		{
			name: "Given calling WithPaginate and give it data object {'limit':10} Should return that" +
				" as paginate object",
			sampleData: obj,
			expect: func(data any) AppSuccess {
				return AppSuccess{Pagination: data}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(AppSuccess)
			WithPaginate(tc.sampleData)(app)
			assert.Equal(t, tc.expect(tc.sampleData), *app)
		})
	}
}

func TestWithErr(t *testing.T) {
	var obj struct {
		Username string `validate:"required"`
	}
	var errValidation = []validationError{{Field: "username", Message: "Param username should not be empty"}}

	testCases := []struct {
		name      string
		sampleErr error
		expect    AppError
	}{
		{
			name: "Given calling WithError and give it error 'oops' should return AppError with code " +
				"'UnknownError' and message 'oops'",
			sampleErr: errors.New("oops"),
			expect:    AppError{Code: "UnknownError", Message: "oops"},
		},
		{
			name: "Given calling WithError and give it Std with Code 'CustomError' Message 'Something is wrong'" +
				" should return exactly same AppError with code 'CustomError' and message 'Something is wrong'",
			sampleErr: NewStd("CustomError", "Something is wrong"),
			expect:    AppError{Code: "CustomError", Message: "Something is wrong"},
		},
		{
			name: "Given calling WithError and give it validation error should return AppError with code " +
				"'ValidationError' and message 'Provided data is invalid, please check again' and data contain" +
				" detailed information about which param is invalid",
			sampleErr: validator.New().Struct(obj),
			expect:    AppError{Code: "ValidationError", Message: "Provided data is invalid, please check again", Error: errValidation},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(AppError)
			WithErr(tc.sampleErr)(app)
			assert.Equal(t, tc.expect, *app)
		})
	}
}
