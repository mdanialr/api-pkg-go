package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStd(t *testing.T) {
	testCases := []struct {
		name       string
		sampleCode string
		sampleMsg  string
		expectObj  *Std
	}{
		{
			name: "Given code 'RecordNotFound' without message should just return pointer to Std with empty " +
				"message field and code 'RecordNotFound'",
			sampleCode: "RecordNotFound",
			expectObj: &Std{
				Code: "RecordNotFound",
			},
		},
		{
			name: "Given code 'InternalError' and message 'Something was wrong!' should return pointer to Std " +
				"complete with both code 'InternalError' and message 'Something was wrong!'",
			sampleCode: "InternalError",
			sampleMsg:  "Something was wrong!",
			expectObj: &Std{
				Code:    "InternalError",
				Message: "Something was wrong!",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := NewStd(tc.sampleCode, tc.sampleMsg)
			assert.Equal(t, tc.expectObj, out)
		})
	}
}

func TestStd_Error(t *testing.T) {
	testCases := []struct {
		name      string
		sample    Std
		expectMsg string
	}{
		{
			name: "Given Std with message 'Something was wrong!' should also return 'Something was wrong!' when " +
				"calling the method Error that's compatible with error interface",
			sample:    Std{Message: "Something was wrong!"},
			expectMsg: "Something was wrong!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectMsg, tc.sample.Error())
		})
	}
}

func TestStd_String(t *testing.T) {
	testCases := []struct {
		name      string
		sample    Std
		expectMsg string
	}{
		{
			name: "Given Std with message 'Something was wrong!' and code 'PanicError' should also return " +
				"'PanicError - Something was wrong!' when calling the method String that's compatible with " +
				"Stringer interface",
			sample:    Std{Message: "Something was wrong!", Code: "PanicError"},
			expectMsg: "PanicError - Something was wrong!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectMsg, tc.sample.String())
		})
	}
}
