// Package logger provide any app-logger related utility.
package logger

import (
	"sync"
	"time"

	"github.com/mdanialr/api-pkg-go/logger"
	"github.com/mdanialr/api-pkg-go/logger/zap/output"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once
var zapLogger *zap.Logger
var zapErr error

// NewLogger return ready to use zap logger from given config.
func NewLogger(config *logger.Config) (*zap.Logger, error) {
	// make sure this setup only executed once until the program exited
	once.Do(func() {
		var cores []zapcore.Core
		for _, cr := range config.Cnf.GetStringSlice("log.output") {
			switch cr {
			case "console":
				logLevel := setupZapLogLevelConsole(config.Cnf)
				cores = append(cores, loggeroutput.NewConsoleZapCore(logLevel))
			case "file":
				logLevel := setupZapLogLevelFile(config.Cnf)
				cores = append(cores, loggeroutput.NewFileZapCore(config.Cnf, logLevel))
			case "newrelic":
				nr, err := loggeroutput.NewNRApp(config)
				if err != nil {
					zapErr = err
					continue // skip this loop and do not append it to cores
				}
				logLevel := setupZapLogLevelNR(config.Cnf)
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
