package response

// NewStdErr init new Std by given code and message as the value.
func NewStdErr(code string, err error) error {
	return &Std{Code: code, Message: err.Error()}
}

// NewStd init new Std by given code and message as the value.
func NewStd(code, msg string) error {
	return &Std{Code: code, Message: msg}
}

// Std error implementer that may be used when need to throw just code and
// message as response inside `error` field.
type Std struct {
	Code    string
	Message string
}

// Error method that implement error interface.
func (e Std) Error() string {
	return e.Code + " - " + e.Message
}

// String method that implement Stringer interface.
func (e Std) String() string {
	return e.Code + " - " + e.Message
}

// AppOpt an option signature for App response.
type AppOpt func(*App)

// App standard success response that may be used in every response
// for all handlers.
type App struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      any    `json:"error,omitempty"`
	Pagination any    `json:"pagination,omitempty"` // Pagination additional field when need to serve many Data and want to show pagination info.
}
