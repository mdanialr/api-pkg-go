// Package logger provide any app-logger related utility.
package logger

import (
	"sync"
	"time"

	"github.com/mdanialr/api-pkg-go/logger/zap/output"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once
var zapLogger *zap.Logger
var zapErr error

// NewLogger return ready to use zap logger from given config.
func NewLogger(cnf *viper.Viper) (*zap.Logger, error) {
	// make sure this setup only executed once until the program exited
	once.Do(func() {
		var cores []zapcore.Core
		for _, cr := range cnf.GetStringSlice("log.output") {
			switch cr {
			case "console":
				logLevel := setupZapLogLevelConsole(cnf)
				cores = append(cores, loggeroutput.NewConsoleZapCore(logLevel))
			case "file":
				logLevel := setupZapLogLevelFile(cnf)
				cores = append(cores, loggeroutput.NewFileZapCore(cnf, logLevel))
			case "newrelic":
				nr, err := loggeroutput.NewNRApp(cnf)
				if err != nil {
					zapErr = err
					continue // skip this loop and do not append it to cores
				}
				logLevel := setupZapLogLevelNR(cnf)
				cores = append(cores, loggeroutput.NewNRZapCore(nr, logLevel))
			}
		}
		// merge all chosen cores based on config
		core := zapcore.NewTee(cores...)
		// use sampler. Refer to https://github.com/uber-go/zap/blob/master/FAQ.md#why-sample-application-logs
		core = zapcore.NewSamplerWithOptions(core, time.Second, 8, 2)
		// put to logger in this package, so it can be accessed through context
		zapLogger = zap.New(core)
	})
	return zapLogger, zapErr
}

func setupZapLogLevelConsole(cnf *viper.Viper) zapcore.Level {
	return setupZapLogLevel(cnf.GetString("log.console.level"))
}

func setupZapLogLevelFile(cnf *viper.Viper) zapcore.Level {
	return setupZapLogLevel(cnf.GetString("log.file.level"))
}

func setupZapLogLevelNR(cnf *viper.Viper) zapcore.Level {
	return setupZapLogLevel(cnf.GetString("log.newrelic.level"))
}

func setupZapLogLevel(level string) zapcore.Level {
	lvl := zapcore.DebugLevel
	switch level {
	case "debug":
		lvl = zapcore.DebugLevel
	case "info":
		lvl = zapcore.InfoLevel
	case "warning":
		lvl = zapcore.WarnLevel
	case "error":
		lvl = zapcore.ErrorLevel
	}
	return lvl
}
