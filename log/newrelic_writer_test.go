package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNewrelicWriter(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		msg := "failed to init newrelic app: license length is not 40"
		require.PanicsWithError(t, msg, func() {
			NewNewrelicWriter(DebugLevel, nil)
		})
	})
	t.Run("Should return the expected value in each Writer implementation", func(t *testing.T) {
		cfg := &Config{NR: NRConfig{
			Name:    "name",
			License: "justarandomstringswithfourtylenghtcharss",
		}}
		wr := NewNewrelicWriter(WarnLevel, cfg)
		assert.IsType(t, &newrelicOutput{}, wr.Writer())
		assert.Equal(t, NEWRELIC, wr.Output())
		assert.Equal(t, WarnLevel, wr.Level())

		// just run
		wr.Writer().Write([]byte("message"))
		wr.Wait(-1)
		wr.Flush(-1)
	})
}
