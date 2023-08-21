package loggeroutput

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFileZerologWriter return new io.Writer that may be used by zerolog writer
// to write json log message to designated local file. The log file will use
// log rotate provided by lumberjack pkg and set it according to the value in
// given config. Set the log level by given param otherwise will use Debug
// level.
func NewFileZerologWriter(cnf *viper.Viper, lvl ...zerolog.Level) (io.Writer, func(time.Duration)) {
	level := zerolog.DebugLevel
	if len(lvl) > 0 {
		level = lvl[0]
	}

	lr := setupFileZerologRotate(cnf)
	fl := &zerologCustomFileWriter{lr, level}
	return fl, fl.Flush
}

type zerologCustomFileWriter struct {
	io.WriteCloser
	level zerolog.Level
}

func (z *zerologCustomFileWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= z.level {
		return z.WriteCloser.Write(p)
	}
	return len(p), nil
}

func (z *zerologCustomFileWriter) Flush(_ time.Duration) {
	z.Close()
}

// setupFileZerologRotate init and set default value to lumberjack.Logger if no
// value provided in given config.
func setupFileZerologRotate(cnf *viper.Viper) *lumberjack.Logger {
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
