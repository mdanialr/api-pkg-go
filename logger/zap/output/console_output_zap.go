package loggeroutput

import (
	"os"

	"github.com/mdanialr/api-pkg-go/logger/zap/output/option"

	"go.uber.org/zap/zapcore"
)

// NewConsoleZapCore return new zapcore.Core that write logs to stdout and use
// zap console encoder. Set the log level by given param otherwise will use
// Debug level. This is best to set in local env.
func NewConsoleZapCore(lvl ...zapcore.Level) zapcore.Core {
	level := zapcore.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	encConfig := loggeroption.NewZapEncConf(
		loggeroption.WithEncStdTimeFmt(),
		loggeroption.WithEncUpperCasedLevel(),
	)
	enc := zapcore.NewConsoleEncoder(encConfig)
	return zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), level)
}
