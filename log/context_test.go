package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCtx(t *testing.T) {
	t.Run("Should return exactly same context if given Producer is nil", func(t *testing.T) {
		expCtx := context.Background()
		ctx := WithCtx(expCtx, nil)
		assert.Equal(t, expCtx, ctx)
	})
	t.Run("Should equal with given Log if singleton instance not assigned yet", func(t *testing.T) {
		pd := NewNop()
		ctx := WithCtx(context.Background(), pd)
		if pr, ok := ctx.Value(ctxKey{}).(Log); ok {
			assert.Equal(t, pr, pd)
		}
	})
	t.Run("Should not equal with given Log if singleton instance already set", func(t *testing.T) {
		singletonLogger = NewNop()
		ctx := WithCtx(context.Background(), nopLog{})
		if pr, ok := ctx.Value(ctxKey{}).(Log); ok {
			assert.NotEqual(t, pr, singletonLogger)
		}
	})
	t.Run("Should be equal if given Log and singleton instance is same", func(t *testing.T) {
		pd := NewNop()
		singletonLogger = pd
		parentCtx := WithCtx(context.Background(), pd)

		childCtx := WithCtx(parentCtx, pd)
		assert.Equal(t, parentCtx, childCtx)
	})
}

func TestFromCtx(t *testing.T) {
	t.Run("Should return no-op Log if no Log in given context", func(t *testing.T) {
		l := FromCtx(context.Background())
		assert.NotNil(t, l)
		assert.Equal(t, &nopLog{}, l)
		assert.IsType(t, &nopLog{}, l)
	})
	t.Run("Should return given Log inside context if singleton is not set", func(t *testing.T) {
		singletonLogger = nil
		nl := NewNop()
		ctx := WithCtx(context.Background(), nl)
		l := FromCtx(ctx)
		assert.NotNil(t, l)
		assert.Equal(t, nl, l)
		assert.IsType(t, &nopLog{}, l)
	})
	t.Run("Should return singleton Log inside context if exist and singleton already set", func(t *testing.T) {
		sl := NewSlogLogger()
		ctx := WithCtx(context.Background(), sl)
		l := FromCtx(ctx)
		assert.NotNil(t, l)
		assert.Equal(t, sl, l)
		assert.IsType(t, &slogLog{}, l)
	})
}
