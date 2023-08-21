package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type ctxKey struct{}

// FromCtx return the Logger associated with given ctx. If no logger is
// associated, the default logger is returned, unless it is nil in which case a
// disabled logger is returned.
func FromCtx(ctx context.Context) *zerolog.Logger {
	if z, ok := ctx.Value(ctxKey{}).(*zerolog.Logger); ok {
		return z
	}
	if zeroLogger != nil {
		return zeroLogger
	}
	z := zerolog.Nop()
	return &z
}

// WithCtx return a copy of ctx with given Logger attached.
func WithCtx(ctx context.Context, z *zerolog.Logger) context.Context {
	if l, ok := ctx.Value(ctxKey{}).(*zerolog.Logger); ok {
		if l == z {
			// avoid storing same logger
			return ctx
		}
		if l.GetLevel() == zerolog.Disabled {
			// do not store disabled logger. See https://github.com/rs/zerolog/blob/master/ctx.go#L36
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, z)
}
