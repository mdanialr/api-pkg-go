// Package logger provide any app-logger related utility.
//
// Shamelessly copied from https://github.com/betterstack-community/go-logging/blob/zap/logger/logger.go.
package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey struct{}

// FromCtx return the Logger associated with given ctx. If no logger is
// associated, the default logger is returned, unless it is nil in which case a
// disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	}
	if zapLogger != nil {
		return zapLogger
	}

	return zap.NewNop()
}

// WithCtx return a copy of ctx with given Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			// do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}
