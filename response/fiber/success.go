package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Success return json response with standard success response as the
// structure.
func Success(c *fiber.Ctx, options ...AppOpt) error {
	app := new(App)
	app.Message = "Ok" // set default to 'ok'

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.Status(http.StatusOK).JSON(*app)
}
