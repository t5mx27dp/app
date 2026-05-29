package app

import (
	"context"
	"os"
)

type Option func(*Controller)

func WithContext(ctx context.Context) Option {
	return func(c *Controller) {
		c.ctx, c.cancel = context.WithCancel(ctx)
	}
}

func WithSignal(signal os.Signal) Option {
	return func(c *Controller) {
		c.signals = append(c.signals, signal)

		if c.signal == nil {
			c.signal = make(chan os.Signal, 1)
		}
	}
}
