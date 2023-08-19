package response

import (
	"net/http"

	r "github.com/mdanialr/api-pkg-go/response"

	"github.com/labstack/echo/v4"
)

// Success return json response with standard success response as the
// structure.
func Success(c echo.Context, options ...r.SuccessOption) error {
	app := new(r.AppSuccess)
	app.Message = "Ok" // set default to 'ok'

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.JSON(http.StatusOK, *app)
}
