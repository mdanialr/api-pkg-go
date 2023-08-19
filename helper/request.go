package helper

import (
	"bytes"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
)

// ExtractRequestHeader extract header from given request and return it as json
// encoded string.
func ExtractRequestHeader(r http.Request) map[string]any {
	b, _ := sonic.ConfigFastest.Marshal(r.Header)
	b = bytes.TrimSpace(b)

	var v map[string]any
	sonic.ConfigDefault.Unmarshal(b, &v)

	return v
}

// ExtractRequest encode given request to json and just return {} if the
// request body is empty.
func ExtractRequest(r *http.Request) map[string]any {
	b, _ := io.ReadAll(r.Body)
	if len(b) > 0 {
		b = bytes.TrimSpace(b)
		r.Body = io.NopCloser(bytes.NewReader(b)) // return back the body to http.Request

		var v map[string]any
		sonic.ConfigFastest.Unmarshal(b, &v)

		return v
	}
	return map[string]any{}
}
