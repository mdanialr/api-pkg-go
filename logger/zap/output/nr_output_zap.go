package loggeroutput

import (
	"io"
	"os"
	"strings"

	"github.com/mdanialr/api-pkg-go/logger/zap/output/option"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// NewNRZapCore return new zapcore.Core that write/send logs directly to
// newrelic server using given nr app.
func NewNRZapCore(nr *newrelic.Application, lvl ...zapcore.Level) zapcore.Core {
	level := zapcore.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	encConfig := loggeroption.NewZapEncConf(
		loggeroption.WithEncStdTimeFmt(),
		loggeroption.WithEncUpperCasedLevel(),
		loggeroption.WithEncTimeKey("time"),
	)
	enc := NewNRZapEncoder(zapcore.NewJSONEncoder(encConfig), nr)
	return zapcore.NewCore(enc, zapcore.AddSync(io.Discard), level)
}

// NewNRApp init new relic app using app name and license from given config.
func NewNRApp(cnf *viper.Viper) (*newrelic.Application, error) {
	name := cnf.GetString("log.newrelic.name")
	license := cnf.GetString("log.newrelic.license")

	return newrelic.NewApplication(
		newrelic.ConfigAppName(name),
		newrelic.ConfigLicense(license),
		newrelic.ConfigInfoLogger(os.Stdout),
	)
}

// NewNRZapEncoder custom zap encoder that send log to given newrelic app.
//
// Thanks to uncle stackoverflow finally can correctly capture zap log along
// with the context fields. Ref:
//   - https://stackoverflow.com/questions/70489167/uber-zap-logger-how-to-prepend-every-log-entry-with-a-string/70492904#70492904
//   - https://stackoverflow.com/questions/70512120/how-to-access-fields-in-zap-hooks/70512858#70512858
//   - https://stackoverflow.com/questions/70502026/why-custom-encoding-is-lost-after-calling-logger-with-in-uber-zap/70503730#70503730
func NewNRZapEncoder(enc zapcore.Encoder, app *newrelic.Application) zapcore.Encoder {
	return &zapEncoder{enc, app}
}

type zapEncoder struct {
	zapcore.Encoder
	nr *newrelic.Application
}

// EncodeEntry custom implement that just call the original EncodeEntry then
// intercept the returned buffer and send it to newrelic.
func (z zapEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// run the original EncodeEntry then grab the returned buffer
	buf, err := z.Encoder.EncodeEntry(entry, fields)

	// prepare the log data
	data := newrelic.LogData{
		Timestamp: entry.Time.UnixMilli(),
		Severity:  entry.Level.String(),
		Message:   strings.TrimSpace(buf.String()), // since the buffer is already json encoded, just inject it as message
	}
	// ingest log to newrelic app
	z.nr.RecordLog(data)

	return buf, err
}

// Clone re-implement the original Clone from zapcore.Encoder. This is
// necessary in case we use zap structured logging instead of the sugared type.
// Ref:
//   - https://github.com/uber-go/zap/blob/master/zapcore/console_encoder.go#L66
//   - https://stackoverflow.com/questions/70512120/how-to-access-fields-in-zap-hooks/70512858#70512858
func (z zapEncoder) Clone() zapcore.Encoder {
	return zapEncoder{z.Encoder.Clone(), z.nr}
}
