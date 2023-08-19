package loggeroption

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapEncConf return zapcore.EncoderConfig that apply the provided options.
func NewZapEncConf(opts ...ZapOption) zapcore.EncoderConfig {
	enc := zap.NewProductionEncoderConfig() // use prod as default encoder config
	// apply options
	for _, opt := range opts {
		opt(&enc)
	}
	return enc
}
