package response

import (
	"net/http"

	r "github.com/mdanialr/api-pkg-go/response"

	"github.com/gofiber/fiber/v2"
)

// Error return fiber framework json response with standard error response as
// the structure.
func Error(c *fiber.Ctx, options ...r.AppErrorOption) error {
	app := new(r.AppError)
	// set default code and message
	app.Code = "UnknownError"
	app.Message = "Something was wrong"

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.Status(http.StatusBadRequest).JSON(*app)
}

// ErrorCode return fiber framework json response with standard error response
// as the structure and use given code as response status code.
func ErrorCode(c *fiber.Ctx, code int, options ...r.AppErrorOption) error {
	app := new(r.AppError)
	// set default code and message
	app.Code = "UnknownError"
	app.Message = "Something was wrong"

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.Status(code).JSON(*app)
}
