package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	r "github.com/mdanialr/api-pkg-go/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	type obj struct {
		User string `validate:"required"`
	}
	var o obj

	testCases := []struct {
		name       string
		sampleErr  error
		sampleOpts func(err error) []r.AppErrorOption
		expectCode int
		expectJson string
	}{
		{
			name: "Given calling Error without any additional options should just return 400 response code " +
				"with json response code UnknownError and message Something was wrong",
			sampleOpts: func(_ error) []r.AppErrorOption {
				return []r.AppErrorOption{}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it error 'oops' should return 400 " +
				"response code with json response code 'UnknownError' and message 'oops'",
			sampleErr: errors.New("oops"),
			sampleOpts: func(err error) []r.AppErrorOption {
				return []r.AppErrorOption{
					r.WithErr(err),
				}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"UnknownError","message":"oops"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it Std that has code " +
				"'RecordNotFound' and message 'try again' should return 400 response code with json response" +
				"code 'RecordNotFound' and message 'try again'",
			sampleErr: r.NewStd("RecordNotFound", "try again"),
			sampleOpts: func(err error) []r.AppErrorOption {
				return []r.AppErrorOption{
					r.WithErr(err),
				}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"RecordNotFound","message":"try again"}`,
		},
		{
			name: "Given calling Error with an option WithError and give it ValidationError that has an " +
				"error 'required' for field 'user' should return 400 response code with json response code" +
				"'ValidationError' and message 'Provided data is invalid, please check again'" +
				" also the validation error message 'required' which is 'Param user should not be empty'",
			sampleErr: validator.New().Struct(o),
			sampleOpts: func(err error) []r.AppErrorOption {
				return []r.AppErrorOption{
					r.WithErr(err),
				}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"ValidationError","message":"Provided data is invalid, please check again","error":[{"field":"user","message":"Param user should not be empty"}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			err := Error(c, tc.sampleOpts(tc.sampleErr)...)

			if assert.NoError(t, err) {
				assert.Equal(t, tc.expectCode, rec.Code)
				assert.Equal(t, tc.expectJson, strings.TrimSpace(rec.Body.String())) // remove default linebreak from json encoder
			}
		})
	}
}

func TestErrorCode(t *testing.T) {
	testCases := []struct {
		name       string
		sampleCode int
		sampleErr  error
		sampleOpts func(err error) []r.AppErrorOption
		expectCode int
		expectJson string
	}{
		{
			name: "Given calling Error Code with 400 without any additional options should just return 400 response code " +
				"with json response code UnknownError and message Something was wrong",
			sampleCode: http.StatusBadRequest,
			sampleOpts: func(_ error) []r.AppErrorOption {
				return []r.AppErrorOption{}
			},
			expectCode: http.StatusBadRequest,
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error Code with 500 without any additional options should just return 500 response code " +
				"with json response code UnknownError and message Something was wrong",
			sampleCode: http.StatusInternalServerError,
			sampleOpts: func(_ error) []r.AppErrorOption {
				return []r.AppErrorOption{}
			},
			expectCode: http.StatusInternalServerError,
			expectJson: `{"code":"UnknownError","message":"Something was wrong"}`,
		},
		{
			name: "Given calling Error Code with 404 and with an option WithError also give it Std that has code " +
				"'RecordNotFound' and message 'not found' should return 404 response code with json response" +
				"code 'RecordNotFound' and message 'not found'",
			sampleErr:  r.NewStd("RecordNotFound", "not found"),
			sampleCode: http.StatusNotFound,
			sampleOpts: func(err error) []r.AppErrorOption {
				return []r.AppErrorOption{
					r.WithErr(err),
				}
			},
			expectCode: http.StatusNotFound,
			expectJson: `{"code":"RecordNotFound","message":"not found"}`,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		err := ErrorCode(c, tc.sampleCode, tc.sampleOpts(tc.sampleErr)...)

		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectCode, rec.Code)
			assert.Equal(t, tc.expectJson, strings.TrimSpace(rec.Body.String())) // remove default linebreak from json encoder
		}
	}
}
