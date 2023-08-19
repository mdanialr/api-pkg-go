package helper

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractRequestHeader(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	req.Header.Add("hello", "world")
	req2 := req.Clone(context.TODO()) // clone so the memory address also differ
	req2.Header.Add("hello", "panda")

	testCases := []struct {
		name    string
		sample  http.Request
		expect  map[string]any
		wantEmp bool
	}{
		{
			name:    "Given empty request header should return empty map",
			sample:  http.Request{},
			expect:  map[string]any{},
			wantEmp: true,
		},
		{
			name:   "Given non-empty request header should return the map version of that header",
			sample: *req,
			expect: map[string]any{
				"Hello": []any{"world"},
			},
		},
		{
			name: "Given non-empty request header should return the map version of that header and if it has " +
				"multi values then should be returned as slice of any type with one key",
			sample: *req2,
			expect: map[string]any{
				"Hello": []any{"world", "panda"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ExtractRequestHeader(tc.sample)
			if tc.wantEmp {
				assert.Empty(t, res)
				return
			}
			assert.NotEmpty(t, res)
			assert.Equal(t, res, tc.expect)
		})
	}
}

func TestExtractRequest(t *testing.T) {
	reqEmpty, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(``)))
	reqBody, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"hello":"world"}`))

	testCases := []struct {
		name    string
		sample  *http.Request
		expect  map[string]any
		wantEmp bool
	}{
		{
			name:    "Given empty request body should return empty map",
			sample:  reqEmpty,
			expect:  map[string]any{},
			wantEmp: true,
		},
		{
			name:   "Given non-empty request body should return the map version of that body payload",
			sample: reqBody,
			expect: map[string]any{
				"hello": "world",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ExtractRequest(tc.sample)
			if tc.wantEmp {
				assert.Empty(t, res)
				return
			}
			assert.NotEmpty(t, res)
			assert.Equal(t, res, tc.expect)
		})
	}
}
