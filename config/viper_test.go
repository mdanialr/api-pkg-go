package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfigYml(t *testing.T) {
	testCases := []struct {
		name       string
		sampleName string
		wantErr    bool
	}{
		{
			name:       "Given config file name that does not exist should return error",
			sampleName: "non-exist",
			wantErr:    true,
		},
		{
			name:       "Given config file name 'config' and does exist should not return error",
			sampleName: "config",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := setupConfig()
			defer os.Remove(f.Name())
			vp, err := InitConfigYml(tc.sampleName)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, "world", vp.Get("hello"))
		})
	}
}

func setupConfig() *os.File {
	const tempViper = `hello: world`
	f, _ := os.Create("config.yml")
	f.WriteString(tempViper)
	f.Close()
	return f
}
