package response

import (
	"net/http"

	r "github.com/mdanialr/api-pkg-go/response"

	"github.com/labstack/echo/v4"
)

// Error return echo framework json response with standard error response as
// the structure.
func Error(c echo.Context, options ...r.AppErrorOption) error {
	app := new(r.AppError)
	// set default code and message
	app.Code = "UnknownError"
	app.Message = "Something was wrong"

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.JSON(http.StatusBadRequest, *app)
}

// ErrorCode return echo framework json response with standard error response
// as the structure and use given code as response status code.
func ErrorCode(c echo.Context, code int, options ...r.AppErrorOption) error {
	app := new(r.AppError)
	// set default code and message
	app.Code = "UnknownError"
	app.Message = "Something was wrong"

	// apply all available options
	for _, opt := range options {
		opt(app)
	}

	return c.JSON(code, *app)
}
