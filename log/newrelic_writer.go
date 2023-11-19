package log

import (
	"bytes"
	"errors"
	"io"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// NewNewrelicWriter return Writer implementer that ingest logs directly to
// newrelic server by given Config.NR and set given Level as the log level.
func NewNewrelicWriter(lvl Level, cnf *Config) Writer {
	if cnf == nil {
		cnf = &Config{}
	}

	nr, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cnf.NR.Name),
		newrelic.ConfigLicense(cnf.NR.License),
		newrelic.ConfigInfoLogger(os.Stdout),
	)
	if err != nil {
		err = errors.New("failed to init newrelic app: " + err.Error())
		panic(err)
	}
	return &newrelicOutput{lvl: lvl, nr: nr}
}

type newrelicOutput struct {
	nr  *newrelic.Application
	lvl Level
}

// Write implement io.Writer.
func (n *newrelicOutput) Write(p []byte) (_ int, err error) {
	msg := string(bytes.TrimSpace(p))
	n.nr.RecordLog(newrelic.LogData{Message: msg})
	return len(p), nil
}
func (n *newrelicOutput) Writer() io.Writer       { return n }
func (n *newrelicOutput) Output() Output          { return NEWRELIC }
func (n *newrelicOutput) Level() Level            { return n.lvl }
func (n *newrelicOutput) Wait(dur time.Duration)  { n.nr.WaitForConnection(dur) }
func (n *newrelicOutput) Flush(dur time.Duration) { n.nr.Shutdown(dur) }
