package response

// NewStd init new Std by given code and message as the value.
func NewStd(code, msg string) *Std {
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
	return e.Message
}

// String method that implement Stringer interface.
func (e Std) String() string {
	return e.Code + " - " + e.Message
}

// AppSuccessOption an option for AppSuccess response.
type AppSuccessOption func(*AppSuccess)

// AppSuccess standard success response that may be used in every response
// for all handlers.
type AppSuccess struct {
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"` // Pagination additional field when need to serve many Data and want to show pagination info.
}

// AppErrorOption an option for AppError response.
type AppErrorOption func(*AppError)

// AppError standard error response that may be used in every response for
// all handlers.
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}
