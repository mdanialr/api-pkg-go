package logger

import "time"

// Log unified front-end to logger.
type Log interface {
	// Init do initialization and should be called at very first before other
	// functions.
	Init(dur time.Duration)
	// Flush do clean up task that make sure any leftover logs written properly
	// to the destination by Log therefor should be called at very last right
	// before program exit for example.
	Flush(dur time.Duration)
	// With add given Log(s) as structured context.
	With(pr ...LogObj) Log
	// Dbg logs a message at DebugLevel.
	Dbg(msg string, pr ...LogObj)
	// Inf logs a message at InfoLevel.
	Inf(msg string, pr ...LogObj)
	// Wrn logs a message at WarnLevel.
	Wrn(msg string, pr ...LogObj)
	// Err logs a message at ErrorLevel.
	Err(msg string, pr ...LogObj)
}

// NewNop returns a no-op Log. Do nothing and never writes out any logs.
func NewNop() Log {
	return &nopLog{}
}

type nopLog struct{}

func (n nopLog) Init(_ time.Duration)      {}
func (n nopLog) Flush(_ time.Duration)     {}
func (n nopLog) With(_ ...LogObj) Log      { return nil }
func (n nopLog) Dbg(_ string, _ ...LogObj) {}
func (n nopLog) Inf(_ string, _ ...LogObj) {}
func (n nopLog) Wrn(_ string, _ ...LogObj) {}
func (n nopLog) Err(_ string, _ ...LogObj) {}
