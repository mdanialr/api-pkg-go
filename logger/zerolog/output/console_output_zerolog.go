package loggeroutput

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// NewConsoleZerologWriter return new io.Writer that may be used by zerolog
// writer to write prettified log message to console. Set the log level by
// given param otherwise will use Debug level. This is best to set in local
// env.
func NewConsoleZerologWriter(lvl ...zerolog.Level) io.Writer {
	level := zerolog.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	console := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stdout
	})

	return &zerologCustomConsoleWriter{console, level}
}

type zerologCustomConsoleWriter struct {
	zerolog.ConsoleWriter
	level zerolog.Level
}

func (z *zerologCustomConsoleWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= z.level {
		return z.ConsoleWriter.Write(p)
	}
	return len(p), nil
}
