package logger

import (
	"io"
	"sync"
	"time"

	"github.com/mdanialr/api-pkg-go/logger"
	loggeroutput "github.com/mdanialr/api-pkg-go/logger/zerolog/output"

	"github.com/rs/zerolog"
)

var once sync.Once
var zeroLogger *zerolog.Logger
var zeroLogWait func(time.Duration)
var zeroLogFlush func(time.Duration)
var zeroLogErr error

// NewLogger return ready to use zerolog logger from given config.
func NewLogger(config *logger.Config) (*zerolog.Logger, func(time.Duration), func(time.Duration), error) {
	once.Do(func() {
		var writers []io.Writer
		var flushes []func(time.Duration)
		for _, cr := range config.Cnf.GetStringSlice("log.output") {
			switch cr {
			case "console":
				logLevel := setupZeroLogLevelConsole(config)
				writers = append(writers, loggeroutput.NewConsoleZerologWriter(logLevel))
			case "file":
				logLevel := setupZeroLogLevelFile(config)
				flLog, flush := loggeroutput.NewFileZerologWriter(config.Cnf, logLevel)
				writers = append(writers, flLog)
				flushes = append(flushes, flush)
			case "newrelic":
				nr, err := loggeroutput.NewNRApp(config)
				if err != nil {
					zeroLogErr = err
					continue // skip this loop and do not append it to writers
				}
				zeroLogWait = func(dur time.Duration) {
					nr.WaitForConnection(dur)
				}
				logLevel := setupZeroLogLevelNR(config)
				nrLog, flush := loggeroutput.NewNRZerologWriter(nr, logLevel)
				writers = append(writers, nrLog)
				flushes = append(flushes, flush)
			}
		}

		// wrap up all chosen writers
		writer := zerolog.MultiLevelWriter(writers...)
		zl := zerolog.New(writer).Level(zerolog.DebugLevel)
		zeroLogger = &zl
		// wrap up all clean up task
		zeroLogFlush = func(dur time.Duration) {
			for _, flush := range flushes {
				flush(dur)
			}
		}
	})
	return zeroLogger, zeroLogWait, zeroLogFlush, zeroLogErr
}

func setupZeroLogLevelConsole(cnf *logger.Config) zerolog.Level {
	return setupZeroLogLevel(cnf.ConsoleLevel)
}

func setupZeroLogLevelFile(cnf *logger.Config) zerolog.Level {
	return setupZeroLogLevel(cnf.FileLevel)
}

func setupZeroLogLevelNR(cnf *logger.Config) zerolog.Level {
	return setupZeroLogLevel(cnf.NRLevel)
}

func setupZeroLogLevel(level string) zerolog.Level {
	lvl := zerolog.DebugLevel
	switch level {
	case "debug":
		lvl = zerolog.DebugLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "warning":
		lvl = zerolog.WarnLevel
	case "error":
		lvl = zerolog.ErrorLevel
	}
	return lvl
}
