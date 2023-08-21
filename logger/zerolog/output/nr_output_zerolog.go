package loggeroutput

import (
	"io"
	"os"
	"time"

	"github.com/mdanialr/api-pkg-go/logger"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

// NewNRZerologWriter return new io.Writer that may be used by zerolog writer
// to write json log message to NewRelic API. Set the log level by given param
// otherwise will use Debug level.
func NewNRZerologWriter(nr *newrelic.Application, lvl ...zerolog.Level) (io.Writer, func(time.Duration)) {
	level := zerolog.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	nrWr := &zerologCustomNRWriter{nr, level}
	return nrWr, nrWr.nr.Shutdown
}

// NewNRApp init new relic app using app name and license from given config.
func NewNRApp(cnf *logger.Config) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(cnf.NRApp),
		newrelic.ConfigLicense(cnf.NRLicense),
		newrelic.ConfigInfoLogger(os.Stdout),
	)
}

type zerologCustomNRWriter struct {
	nr    *newrelic.Application
	level zerolog.Level
}

func (z *zerologCustomNRWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (z *zerologCustomNRWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= z.level {
		record := newrelic.LogData{
			Severity: level.String(),
			Message:  string(p),
		}
		z.nr.RecordLog(record)
	}
	return len(p), nil
}
