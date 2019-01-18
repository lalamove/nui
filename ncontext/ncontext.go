package ncontext

import (
	"context"
	"time"
)

// Contexter is an interfacre to provide contexts
type Contexter interface {
	WithTimeout(context.Context, time.Duration) (context.Context, context.CancelFunc)
	WithDeadline(context.Context, time.Time) (context.Context, context.CancelFunc)
	WithCancel(context.Context) (context.Context, context.CancelFunc)
}

// DefaultContexter is a default implementation of the Contexter interface which wraps the context package
var DefaultContexter = defaultContexter{}

type defaultContexter struct{}

func (d defaultContexter) WithTimeout(ctx context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, duration)
}

func (d defaultContexter) WithDeadline(ctx context.Context, t time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(ctx, t)
}

func (d defaultContexter) WithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
