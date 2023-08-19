package loggeroption

import "go.uber.org/zap/zapcore"

// ZapOption option signature that may be used to optionally add custom setting
// to zap encoder config.
type ZapOption func(cnf *zapcore.EncoderConfig)

// WithEnableCallerFunc use default zapcore.ShortCallerEncoder as the func
// caller encoder.
func WithEnableCallerFunc() ZapOption {
	return func(cnf *zapcore.EncoderConfig) {
		cnf.EncodeCaller = zapcore.ShortCallerEncoder
	}
}

// WithEncStdTimeFmt change encoder config timestamp format to time.RFC3339.
func WithEncStdTimeFmt() ZapOption {
	return func(cnf *zapcore.EncoderConfig) {
		cnf.EncodeTime = zapcore.RFC3339TimeEncoder
	}
}

// WithEncTimeKey change the default encoder config timestamp key 'ts' to given
// string.
func WithEncTimeKey(k string) ZapOption {
	return func(cnf *zapcore.EncoderConfig) {
		cnf.TimeKey = k
	}
}

// WithEncUpperCasedLevel change encoder config level 'info' to 'INFO'.
func WithEncUpperCasedLevel() ZapOption {
	return func(cnf *zapcore.EncoderConfig) {
		cnf.EncodeLevel = zapcore.CapitalLevelEncoder
	}
}
