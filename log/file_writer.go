package log

import (
	"io"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewFileWriter return Writer implementer that write logs to designated file
// based on the given FileConfig and set given Level as the log level.
func NewFileWriter(lvl Level, cnf *Config) Writer {
	if cnf == nil {
		cnf = &Config{}
	}

	return &fileOutputWithLumberjack{lvl: lvl, wr: setupLumberjack(&cnf.File)}
}

type fileOutputWithLumberjack struct {
	wr  *lumberjack.Logger
	lvl Level
}

func (f *fileOutputWithLumberjack) Writer() io.Writer     { return f.wr }
func (f *fileOutputWithLumberjack) Output() Output        { return FILE }
func (f *fileOutputWithLumberjack) Level() Level          { return f.lvl }
func (f *fileOutputWithLumberjack) Wait(_ time.Duration)  {}
func (f *fileOutputWithLumberjack) Flush(_ time.Duration) { f.wr.Close() }

// setupLumberjack init and set default value to lumberjack.Logger if no value
// provided in given config.
func setupLumberjack(cnf *FileConfig) *lumberjack.Logger {
	lj := lumberjack.Logger{
		Filename:   cnf.Path,
		MaxSize:    cnf.Size,
		MaxAge:     cnf.Age,
		MaxBackups: cnf.Num,
		LocalTime:  true,
	}

	// set default value
	if lj.Filename == "" {
		lj.Filename = "./logs/app.log"
	}
	if lj.MaxSize == 0 {
		lj.MaxSize = 150
	}
	if lj.MaxAge == 0 {
		lj.MaxAge = 28
	}
	if lj.MaxBackups == 0 {
		lj.MaxBackups = 7
	}
	return &lj
}
