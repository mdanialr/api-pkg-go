package response

import (
	"testing"

	help "github.com/mdanialr/api-pkg-go/helper"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestNewValidationErrors(t *testing.T) {
	testCases := []struct {
		name        string
		sampleErr   func() validator.ValidationErrors
		expectField string
		expectMsg   string
	}{
		{
			name: "Given calling NewValidationErrors with an error from `required` tag and field name `user` " +
				"should return slice of struct validation error with Field as 'user' and Message as 'Param" +
				" user should not be empty'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					User string `validate:"required"`
				}
				obj.User = ""
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "user",
			expectMsg:   "Param user should not be empty",
		},
		{
			name: "Given calling NewValidationErrors with an error from `len=5` tag and field name `username` " +
				"should return slice of struct validation error with Field as 'username' and Message as 'Param" +
				" username should have length 5 characters'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"len=5"`
				}
				obj.Username = "1234"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Param username should have length 5 characters",
		},
		{
			name: "Given calling NewValidationErrors with an error from `numeric` tag and field name `code` " +
				"should return slice of struct validation error with Field as 'code' and Message as 'Param" +
				" code should only contain numeric characters'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Code string `validate:"numeric"`
				}
				obj.Code = "abc"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "code",
			expectMsg:   "Param code should only contain numeric characters",
		},
		{
			name: "Given calling NewValidationErrors with an error from `max=10` tag and field name `username` " +
				"should return slice of struct validation error with Field as 'username' and Message as 'Param" +
				" username should have value or length characters same or less than 10'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"max=10"`
				}
				obj.Username = "Lorem Ipsum is simply dummy text of the printing and typesetting industry"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Param username should have value or length characters same or less than 10",
		},
		{
			name: "Given calling NewValidationErrors with an error from `min=10` tag and field name `username` " +
				"should return slice of struct validation error with Field as 'username' and Message as 'Param" +
				" username should have value or length characters same or more than 10'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"min=10"`
				}
				obj.Username = "123456789"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Param username should have value or length characters same or more than 10",
		},
		{
			name: "Given calling NewValidationErrors with an error from `url` tag and field name `image` " +
				"should return slice of struct validation error with Field as 'image' and Message as 'Param" +
				" image should be valid url and have FQDN'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Image string `validate:"url"`
				}
				obj.Image = "example.com/image.png"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "image",
			expectMsg:   "Param image should be valid url and have FQDN",
		},
		{
			name: "Given calling NewValidationErrors with an error from `boolean` tag and field name `status` " +
				"should return slice of struct validation error with Field as 'status' and Message as 'Param" +
				" status should only contain either (true | false)'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Status string `validate:"boolean"`
				}
				obj.Status = "yes"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "status",
			expectMsg:   "Param status should only contain either (true | false)",
		},
		{
			name: "Given calling NewValidationErrors with an error from `alpha` tag and field name `username` " +
				"should return slice of struct validation error with Field as 'username' and Message as 'Param" +
				" username should only contain alphabet characters'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"alpha"`
				}
				obj.Username = "123"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Param username should only contain alphabet characters",
		},
		{
			name: "Given calling NewValidationErrors with an error from `alphanum` tag and field name `username` " +
				"should return slice of struct validation error with Field as 'username' and Message as 'Param" +
				" username should only contain alphabet and or numeric characters'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"alphanum"`
				}
				obj.Username = "123fg!"
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Param username should only contain alphabet and or numeric characters",
		},
		{
			name: "Given calling NewValidationErrors with an error from `image` tag and field name `image` " +
				"should return slice of struct validation error with Field as 'image' and Message as 'Param" +
				" image should only contain one of (jpg|jpeg|png) file extension'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Image string `validate:"image"`
				}
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "image",
			expectMsg:   "Param image should only contain one of (jpg|jpeg|png) file extension",
		},
		{
			name: "Given calling NewValidationErrors with an error from `gte=18` tag and field name `age` " +
				"should return slice of struct validation error with Field as 'age' and Message as 'Param" +
				" age should have value same or more than 18'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Age int `validate:"gte=18"`
				}
				obj.Age = 17
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "age",
			expectMsg:   "Param age should have value same or more than 18",
		},
		{
			name: "Given calling NewValidationErrors with an error from `ltefield=MinCap` tag and field name `age` " +
				"should return slice of struct validation error with Field as 'age' and Message as 'Param" +
				" age should have value that's more than the value of param min_cap'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					MinCap int
					Age    int `validate:"ltefield=MinCap"`
				}
				obj.MinCap = 17
				obj.Age = 18
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "age",
			expectMsg:   "Param age should have value that's more than the value of param min_cap",
		},
		{
			name: "Given calling NewValidationErrors with an error from `ltfield=MinCap` tag and field name `age` " +
				"should return slice of struct validation error with Field as 'age' and Message as 'Param" +
				" age should have value that's equal or more than the value of param min_cap'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					MinCap int
					Age    int `validate:"ltfield=MinCap"`
				}
				obj.MinCap = 17
				obj.Age = 17
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "age",
			expectMsg:   "Param age should have value that's equal or more than the value of param min_cap",
		},
		{
			name: "Given calling NewValidationErrors with an error from `gtfield=MaxCap` tag and field name `age` " +
				"should return slice of struct validation error with Field as 'age' and Message as 'Param" +
				" age should have value that's equal or less than the value of param max_cap'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					MaxCap int
					Age    int `validate:"gtfield=MaxCap"`
				}
				obj.MaxCap = 17
				obj.Age = 16
				return validator.New().Struct(obj).(validator.ValidationErrors)
			},
			expectField: "age",
			expectMsg:   "Param age should have value that's equal or less than the value of param max_cap",
		},
		{
			name: "Given calling NewValidationErrors with an unmapped error such as `unknown` tag and field" +
				" name `age` should return slice of struct validation error with Field as 'age' and default " +
				"validator Message 'Key: 'Age' Error:Field validation for 'Age' failed on the 'unknown' tag'",
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Age string `validate:"unknown"`
				}
				v := validator.New()
				v.RegisterValidation("unknown", func(fl validator.FieldLevel) bool {
					return false
				})

				return v.Struct(obj).(validator.ValidationErrors)
			},
			expectField: "age",
			expectMsg:   "Key: 'Age' Error:Field validation for 'Age' failed on the 'unknown' tag",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := NewValidationErrors(tc.sampleErr())
			assert.Len(t, errors, 1)
			assert.Equal(t, tc.expectField, errors[0].Field)
			assert.Equal(t, tc.expectMsg, errors[0].Message)
		})
	}
}

func TestAddErrorMessage(t *testing.T) {
	testCases := []struct {
		name        string
		sampleKey   string
		sampleVal   func(fe validator.FieldError) string
		sampleErr   func() validator.ValidationErrors
		expectField string
		expectMsg   string
	}{
		{
			name: "Given additional custom error message mapping with key 'custom' and value that return string" +
				" 'Custom Field {field} Validation Error Message' and calling NewValidationErrors with an error from " +
				"`custom` tag and field name `username` should return slice of struct validation error with " +
				"Field as 'username' and Message as 'Custom Field username Validation Error Message'",
			sampleKey: "custom",
			sampleVal: func(fe validator.FieldError) string {
				return "Custom Field " + help.ToSnake(fe.Field()) + " Validation Error Message"
			},
			sampleErr: func() validator.ValidationErrors {
				var obj struct {
					Username string `validate:"custom"`
				}
				v := validator.New()
				v.RegisterValidation("custom", func(fl validator.FieldLevel) bool {
					return false
				})
				return v.Struct(obj).(validator.ValidationErrors)
			},
			expectField: "username",
			expectMsg:   "Custom Field username Validation Error Message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			AddErrorMessage(tc.sampleKey, tc.sampleVal) // add additional error message mapping
			errors := NewValidationErrors(tc.sampleErr())
			assert.Len(t, errors, 1)
			assert.Equal(t, tc.expectField, errors[0].Field)
			assert.Equal(t, tc.expectMsg, errors[0].Message)
		})
	}
}
