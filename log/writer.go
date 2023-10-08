package log

import (
	"io"
	"time"
)

// Writer unified log writer that's responsible where the log from Log
// should be written to.
type Writer interface {
	// Writer return where and how the implementer should write the logs.
	Writer() io.Writer
	// Output define the Output type.
	Output() Output
	// Level define the logs level.
	Level() Level
	// Wait in case the implementer need some delay or preparation before can write any logs.
	Wait(dur time.Duration)
	// Flush any necessary clean up task that will be run by Producer at the last order.
	Flush(dur time.Duration)
}
