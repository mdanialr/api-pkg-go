package loggeroutput

import (
	"github.com/mdanialr/api-pkg-go/logger/zap/output/option"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFileZapCore return new zapcore.Core that write logs to file and use zap
// json encoder. The log file will use log rotate provided by lumberjack pkg
// and set it according to the value in given config.
func NewFileZapCore(cnf *viper.Viper, lvl ...zapcore.Level) zapcore.Core {
	level := zapcore.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	// use lumberjack as log rotation for file output
	w := setupFileZapRotate(cnf)

	// prepare encoder and it's config
	encConf := loggeroption.NewZapEncConf(
		loggeroption.WithEncStdTimeFmt(),
		loggeroption.WithEncUpperCasedLevel(),
		loggeroption.WithEncTimeKey("time"),
	)
	enc := zapcore.NewJSONEncoder(encConf)
	return zapcore.NewCore(enc, zapcore.AddSync(w), level)
}

// setupFileZapRotate set default value to lumberjack.Logger if no value
// provided in given config.
func setupFileZapRotate(cnf *viper.Viper) *lumberjack.Logger {
	lj := lumberjack.Logger{
		Filename:   cnf.GetString("log.file.path"),
		MaxSize:    cnf.GetInt("log.file.size"),
		MaxAge:     cnf.GetInt("log.file.age"),
		MaxBackups: cnf.GetInt("log.file.num"),
		LocalTime:  true,
	}

	// set default value
	if lj.Filename == "" {
		lj.Filename = "./logs/app.log"
	}
	if lj.MaxSize == 0 {
		lj.MaxSize = 25
	}
	if lj.MaxAge == 0 {
		lj.MaxAge = 28
	}
	if lj.MaxBackups == 0 {
		lj.MaxBackups = 7
	}
	return &lj
}
