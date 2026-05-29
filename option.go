package app

import (
	"context"
	"os"
)

type Option func(*Controller)

func WithContext(ctx context.Context) Option {
	return func(c *Controller) {
		ctx, cancel := context.WithCancel(ctx)
		c.ctx = ctx
		c.cancel = cancel
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
