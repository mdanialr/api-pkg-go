package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger return Logger implementer that use zap as the backend.
func NewZapLogger(wr ...Writer) Logger {
	singletonLogger = &zapLog{wr: wr}
	return singletonLogger
}

type zapLog struct {
	log *zap.Logger
	wr  []Writer
}

func (z *zapLog) clone() *zapLog {
	c := *z
	return &c
}
func (z *zapLog) Init(dur time.Duration) {
	var cores []zapcore.Core
	for _, w := range z.wr {
		switch w.Output() {
		case CONSOLE:
			encCnf := zap.NewDevelopmentConfig().EncoderConfig
			// give color to console
			encCnf.EncodeLevel = zapcore.CapitalColorLevelEncoder
			enc := zapcore.NewConsoleEncoder(encCnf)
			core := zapcore.NewCore(enc, zapcore.AddSync(w.Writer()), toZapLevel(w.Level()))
			cores = append(cores, core)
		case FILE, NEWRELIC:
			encCnf := zap.NewProductionEncoderConfig()
			encCnf.EncodeTime = zapcore.RFC3339TimeEncoder
			encCnf.EncodeLevel = zapcore.CapitalLevelEncoder
			encCnf.TimeKey = "time"
			enc := zapcore.NewJSONEncoder(encCnf)
			core := zapcore.NewCore(enc, zapcore.AddSync(w.Writer()), toZapLevel(w.Level()))
			cores = append(cores, core)
		}
		w.Wait(dur)
	}
	z.log = zap.New(zapcore.NewTee(cores...))
}
func (z *zapLog) Flush(dur time.Duration) {
	for _, w := range z.wr {
		w.Flush(dur)
	}
	z.log.Sync()
}
func (z *zapLog) With(pr ...Log) Logger {
	if len(pr) == 0 {
		return z
	}
	mutex.Lock()
	defer mutex.Unlock()

	// clone it, so on every With method call does not affect the parent logger
	clone := z.clone()
	clone.log = clone.log.With(toZapFields(pr)...)
	// then reassign to singleton
	singletonLogger = clone

	return clone
}
func (z *zapLog) Dbg(msg string, pr ...Log) {
	if len(pr) > 0 {
		z.log.Debug(msg, toZapFields(pr)...)
		return
	}
	z.log.Debug(msg)
}
func (z *zapLog) Inf(msg string, pr ...Log) {
	if len(pr) > 0 {
		z.log.Info(msg, toZapFields(pr)...)
		return
	}
	z.log.Info(msg)
}
func (z *zapLog) Wrn(msg string, pr ...Log) {
	if len(pr) > 0 {
		z.log.Warn(msg, toZapFields(pr)...)
		return
	}
	z.log.Warn(msg)
}
func (z *zapLog) Err(msg string, pr ...Log) {
	if len(pr) > 0 {
		z.log.Error(msg, toZapFields(pr)...)
		return
	}
	z.log.Error(msg)
}

// toZapLevel transform local log Level to specific pkg log level which is zap.
func toZapLevel(lvl Level) zapcore.Level {
	switch lvl {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	}
	return zapcore.InvalidLevel
}

// toZapFields transform local Produce to specific pkg log field which is zap.
func toZapFields(pr []Log) []zapcore.Field {
	var fields []zapcore.Field
	for _, p := range pr {
		switch p.typ {
		case StringType:
			fields = append(fields, zap.String(p.key, p.str))
		case NumType:
			fields = append(fields, zap.Int(p.key, p.num))
		case FloatType:
			fields = append(fields, zap.Float64(p.key, p.flt))
		case BoolType:
			fields = append(fields, zap.Bool(p.key, p.b))
		case AnyType:
			fields = append(fields, zap.Any(p.key, p.any))
		}
	}
	return fields
}
