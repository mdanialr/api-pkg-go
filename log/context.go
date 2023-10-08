package log

import "context"

// singletonLogger is holder of Logger and intended to be used as singleton.
var singletonLogger Logger

type ctxKey struct{}

// WithCtx return a copy of ctx with given logger attached.
func WithCtx(ctx context.Context, w Logger) context.Context {
	if w == nil {
		return ctx
	}
	if ww, ok := ctx.Value(ctxKey{}).(Logger); ok {
		// do not store exactly same Logger instance
		if ww == singletonLogger {
			return ctx
		}
	}
	return context.WithValue(ctx, ctxKey{}, w)
}

// FromCtx return the singleton Log associated with given ctx. If no logger is
// associated, the default logger is returned, unless it is nil in which case a
// no-op logger is returned.
func FromCtx(ctx context.Context) Logger {
	if ww, ok := ctx.Value(ctxKey{}).(Logger); ok {
		// use singleton if already set
		if singletonLogger != nil {
			return singletonLogger
		}
		// use that different Log instead
		return ww
	}
	return NewNop()
}
