package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToStr(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		sample []byte
		expect string
	}{
		{
			name:   "Given bytes string 'hello' should return 'hello'",
			sample: []byte(`hello`),
			expect: "hello",
		},
		{
			name:   "Given bytes string 'hello world' should return 'hello world'",
			sample: []byte(`hello world`),
			expect: "hello world",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, BytesToStr(tc.sample))
		})
	}
}

func TestStrToBytes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		sample string
		expect []byte
	}{
		{
			name:   "Given string 'hello' should return 'hello' in bytes",
			sample: "hello",
			expect: StrToBytes("hello"),
		},
		{
			name:   "Given string 'hello world' should return 'hello world' in bytes",
			sample: "hello world",
			expect: StrToBytes("hello world"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
