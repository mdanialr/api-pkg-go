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
		expect     func(data any) App
	}{
		{
			name: "Given calling WithData and give it data object {'username':'user'} Should return that" +
				" as data object",
			sampleData: obj,
			expect: func(data any) App {
				return App{Data: data}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(App)
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
		expect     func(pg any) App
	}{
		{
			name: "Given calling WithPaginate and give it data object {'limit':10} Should return that" +
				" as paginate object",
			sampleData: obj,
			expect: func(data any) App {
				return App{Pagination: data}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(App)
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
		expect    App
	}{
		{
			name: "Given calling WithError and give it error 'oops' should return App with code " +
				"'UnknownError' and message 'oops'",
			sampleErr: errors.New("oops"),
			expect:    App{Code: "UnknownError", Message: "oops"},
		},
		{
			name: "Given calling WithError and give it Std with Code 'CustomError' Message 'Something is wrong'" +
				" should return exactly same App with code 'CustomError' and message 'Something is wrong'",
			sampleErr: NewStd("CustomError", "Something is wrong"),
			expect:    App{Code: "CustomError", Message: "Something is wrong"},
		},
		{
			name: "Given calling WithError and give it validation error should return App with code " +
				"'ValidationError' and message 'Provided data is invalid, please check again' and data contain" +
				" detailed information about which param is invalid",
			sampleErr: validator.New().Struct(obj),
			expect:    App{Code: "ValidationError", Message: "Provided data is invalid, please check again", Error: errValidation},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := new(App)
			WithErr(tc.sampleErr)(app)
			assert.Equal(t, tc.expect, *app)
		})
	}
}
