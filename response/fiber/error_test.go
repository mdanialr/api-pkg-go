package response

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	var o struct {
		Username string `validate:"required"`
	}

	testCases := []struct {
		name       string
		sampleErr  error
		sampleOpts func(error) []AppOpt
		expectJson string
	}{
		{
			name: "Given calling Error without any additional options should just return 400 response code " +
				"with json response code 'UnknownError' and message 'Something was wrong'",
			sampleOpts: func(_ error) []AppOpt {
				return []AppOpt{}
			},
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it error 'oops' should return 400 " +
				"response code with json response code 'UnknownError' and message 'oops'",
			sampleErr: errors.New("oops"),
			sampleOpts: func(err error) []AppOpt {
				return []AppOpt{
					WithErr(err),
				}
			},
			expectJson: `{"code":"UnknownError","message":"oops"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it Std that has code " +
				"'RecordNotFound' and message 'try again' should return 400 response code with json response" +
				"code 'RecordNotFound' and message 'try again'",
			sampleErr: NewStd("RecordNotFound", "try again"),
			sampleOpts: func(err error) []AppOpt {
				return []AppOpt{
					WithErr(err),
				}
			},
			expectJson: `{"code":"RecordNotFound","message":"try again"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it ValidationError that has an " +
				"error 'required' for field 'user' should return 400 response code with json response code" +
				"'ValidationError' and message 'Provided data is invalid, please check again'" +
				" also the validation error message 'required' which is 'Param user should not be empty'",
			sampleErr: validator.New().Struct(o),
			sampleOpts: func(err error) []AppOpt {
				return []AppOpt{
					WithErr(err),
				}
			},
			expectJson: `{"code":"ValidationError","message":"Provided data is invalid, please check again","error":[{"field":"username","message":"Param username should not be empty"}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			f := fiber.New()
			f.Post("/", func(c *fiber.Ctx) error {
				return Error(c, tc.sampleOpts(tc.sampleErr)...)
			})
			resp, err := f.Test(req)

			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				bd, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tc.expectJson, string(bd))
			}
		})
	}
}

func TestErrorCode(t *testing.T) {
	testCases := []struct {
		name       string
		sampleCode int
		sampleErr  error
		sampleOpts func(error) []AppOpt
		expectCode int
		expectJson string
	}{
		{
			name: "Given calling Error Code with 400 without any additional options should just return 400 response code " +
				"with json response code UnknownError and message Something was wrong",
			sampleCode: http.StatusBadRequest,
			sampleOpts: func(_ error) []AppOpt {
				return []AppOpt{}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error Code with 500 without any additional options should just return 500 response code " +
				"with json response code 'UnknownError' and message 'Something was wrong'",
			sampleCode: http.StatusInternalServerError,
			sampleOpts: func(_ error) []AppOpt {
				return []AppOpt{}
			},
			expectCode: http.StatusInternalServerError,
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error Code with 404 and with an option WithError also give it Std that has code " +
				"'RecordNotFound' and message 'not found' should return 404 response code with json response" +
				"code 'RecordNotFound' and message 'not found'",
			sampleErr:  NewStd("RecordNotFound", "not found"),
			sampleCode: http.StatusNotFound,
			sampleOpts: func(err error) []AppOpt {
				return []AppOpt{
					WithErr(err),
				}
			},
			expectCode: http.StatusNotFound,
			expectJson: `{"code":"RecordNotFound","message":"not found"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			f := fiber.New()
			f.Post("/", func(c *fiber.Ctx) error {
				return ErrorCode(c, tc.sampleCode, tc.sampleOpts(tc.sampleErr)...)
			})
			resp, err := f.Test(req)

			if assert.NoError(t, err) {
				assert.Equal(t, tc.expectCode, resp.StatusCode)
				bd, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tc.expectJson, string(bd))
			}
		})
	}
}
