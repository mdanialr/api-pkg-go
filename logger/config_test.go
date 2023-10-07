package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("Applying all available options", func(t *testing.T) {
		cnf := NewConfig(
			WithNRAppName("kpm-library"),
			WithNRLicense("license"),
			WithFilePath("/var/log/app.log"),
			WithFileSize(100),
			WithFileAge(7),
			WithFileMaxBackup(7),
		)

		// assert all values
		assert.Equal(t, "kpm-library", cnf.NR.Name)
		assert.Equal(t, "license", cnf.NR.License)
		assert.Equal(t, "/var/log/app.log", cnf.File.Path)
		assert.Equal(t, 100, cnf.File.Size)
		assert.Equal(t, 7, cnf.File.Age)
		assert.Equal(t, 7, cnf.File.Num)
	})
}
