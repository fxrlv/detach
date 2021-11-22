package detach

import (
	"context"
	"time"
)

func Background(ctx context.Context) context.Context {
	return WithCancel(ctx, context.Background())
}

func WithCancel(ctx, cancel context.Context) context.Context {
	return &detached{
		Context: ctx,
		cancel:  cancel,
	}
}

type detached struct {
	context.Context
	cancel context.Context
}

func (c *detached) Deadline() (time.Time, bool) { return c.cancel.Deadline() }
func (c *detached) Done() <-chan struct{}       { return c.cancel.Done() }
func (c *detached) Err() error                  { return c.cancel.Err() }
