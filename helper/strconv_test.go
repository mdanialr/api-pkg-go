package helper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrYNToBool(t *testing.T) {
	testCases := []struct {
		name   string
		sample string
		expect bool
	}{
		{
			name:   "Given string 'yes' should return true",
			sample: "yes",
			expect: true,
		},
		{
			name:   "Given string 'no' should return false",
			sample: "no",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, StrYNToBool(tc.sample))
		})
	}
}

func TestBoolYNToStr(t *testing.T) {
	testCases := []struct {
		name   string
		sample bool
		expect string
	}{
		{
			name:   "Given bool true should return 'yes'",
			sample: true,
			expect: "yes",
		},
		{
			name:   "Given bool false should return 'no'",
			sample: false,
			expect: "no",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, BoolYNToStr(tc.sample))
		})
	}
}

func TestDropBlankSpace(t *testing.T) {
	testCases := []struct {
		name   string
		sample string
		expect string
	}{
		{
			name:   "Given string 'blank space' should return 'blankspace'",
			sample: "blank space",
			expect: "blankspace",
		},
		{
			name:   "Given string 'nospace' should just return as is 'nospace'",
			sample: "nospace",
			expect: "nospace",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := strings.Map(DropBlankSpace, tc.sample)
			assert.Equal(t, tc.expect, str)
		})
	}
}

func TestToSnake(t *testing.T) {
	testCases := []struct {
		name   string
		sample string
		expect string
	}{
		{
			name:   "Given string 'Hello World' should return 'hello_world'",
			sample: "Hello World",
			expect: "hello_world",
		},
		{
			name:   "Given string 'Drink HiLo' should return 'drink_hi_lo'",
			sample: "Drink HiLo",
			expect: "drink_hi_lo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, ToSnake(tc.sample))
		})
	}
}
