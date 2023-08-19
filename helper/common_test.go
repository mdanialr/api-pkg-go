package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	const c1 = "Given string 'hello' should return pointer of string type and has 'hello' value but different memory address"
	t.Run(c1, func(t *testing.T) {
		var str = "hello"
		pt := Ptr(str)
		// assert value
		assert.Equal(t, str, *pt)
		// assert memory address
		assert.NotSame(t, &str, &pt)
	})
	const c2 = "Given uint '11' should return pointer of uint type and has '11' value but different memory address"
	t.Run(c2, func(t *testing.T) {
		var num = uint(11)
		pt := Ptr(num)
		// assert value
		assert.Equal(t, num, *pt)
		// assert memory address
		assert.NotSame(t, &num, &pt)
	})
}

func TestDef(t *testing.T) {
	var sample = "hello"

	testCases := []struct {
		name   string
		sample *string
		expect string
	}{
		{
			name: "Given pointer to type string with value 'hello' should return different memory address " +
				"but with identical value",
			sample: &sample,
			expect: "hello",
		},
		{
			name: "Given nil pointer to type string without any value should return different memory address" +
				" and empty string",
			sample: nil,
			expect: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := Def(tc.sample)
			// assert same value
			assert.Equal(t, tc.expect, out)
			// assert memory address
			assert.NotSame(t, &tc.sample, &out)
		})
	}
}
