package response

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	r "github.com/mdanialr/api-pkg-go/response"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	testCases := []struct {
		name       string
		sampleOpts func(_ any) []r.AppSuccessOption
		expectJson string
	}{
		{
			name: "Given calling Success without any additional options should just return 200 response code " +
				"with json response message Ok",
			sampleOpts: func(_ any) []r.AppSuccessOption {
				return []r.AppSuccessOption{}
			},
			expectJson: `{"message":"Ok"}`,
		},
		{
			name: "Given calling Success with an option WithData which take object Username with value lorem " +
				"should just return 200 response code " +
				"with json response message Ok and data {'username':'lorem'}",
			sampleOpts: func(_ any) []r.AppSuccessOption {
				var obj = make(map[string]string)
				obj["username"] = "lorem"
				return []r.AppSuccessOption{
					r.WithData(obj),
				}
			},
			expectJson: `{"message":"Ok","data":{"username":"lorem"}}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			f := fiber.New()
			f.Post("/", func(c *fiber.Ctx) error {
				return Success(c, tc.sampleOpts(tc.sampleOpts)...)
			})
			resp, err := f.Test(req)

			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
				bd, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tc.expectJson, string(bd))
			}
		})
	}
}
