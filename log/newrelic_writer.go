package log

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// NewNewrelicWriter return Writer implementer that ingest logs directly to
// newrelic server by given NRConfig and set given Level as the log level.
func NewNewrelicWriter(lvl Level, cnf *Config) Writer {
	nr, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cnf.NR.Name),
		newrelic.ConfigLicense(cnf.NR.License),
		newrelic.ConfigInfoLogger(os.Stdout),
	)
	if err != nil {
		log.Fatalln("failed to init newrelic app:", err)
	}
	return &newrelicOutput{lvl: lvl, nr: nr}
}

type newrelicOutput struct {
	nr  *newrelic.Application
	lvl Level
}

// Write implement io.Writer.
func (n *newrelicOutput) Write(p []byte) (_ int, err error) {
	msg := strings.TrimSpace(string(p))
	n.nr.RecordLog(newrelic.LogData{Message: msg})
	return 0, nil
}
func (n *newrelicOutput) Writer() io.Writer       { return n }
func (n *newrelicOutput) Output() Output          { return NEWRELIC }
func (n *newrelicOutput) Level() Level            { return n.lvl }
func (n *newrelicOutput) Wait(dur time.Duration)  { n.nr.WaitForConnection(dur) }
func (n *newrelicOutput) Flush(dur time.Duration) { n.nr.Shutdown(dur) }