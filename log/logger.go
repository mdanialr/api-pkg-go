package log

import "time"

// Logger unified front-end to log.
type Logger interface {
	// Init do initialization and should be called at very first before
	// other functions.
	Init(dur time.Duration)
	// Flush do clean up task that make sure any leftover logs written properly
	// to the destination by Log therefor should be called at very last
	// after other functions.
	Flush(dur time.Duration)
	// With add given Log(s) as structured context.
	With(pr ...Log) Logger
	// Dbg logs a message at DebugLevel.
	Dbg(msg string, pr ...Log)
	// Inf logs a message at InfoLevel.
	Inf(msg string, pr ...Log)
	// Wrn logs a message at WarnLevel.
	Wrn(msg string, pr ...Log)
	// Err logs a message at ErrorLevel.
	Err(msg string, pr ...Log)
}

// NewNop returns a no-op Logger. Do nothing and never writes out any logs.
func NewNop() Logger {
	return &nopLog{}
}

type nopLog struct{}

func (n nopLog) Init(_ time.Duration)   {}
func (n nopLog) Flush(_ time.Duration)  {}
func (n nopLog) With(_ ...Log) Logger   { return nil }
func (n nopLog) Dbg(_ string, _ ...Log) {}
func (n nopLog) Inf(_ string, _ ...Log) {}
func (n nopLog) Wrn(_ string, _ ...Log) {}
func (n nopLog) Err(_ string, _ ...Log) {}
