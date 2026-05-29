package app

import (
	"os"
)

type Option func(*Controller)

func WithSignal(signal os.Signal) Option {
	return func(c *Controller) {
		c.signals = append(c.signals, signal)

		if c.signal == nil {
			c.signal = make(chan os.Signal, 1)
		}
	}
}
