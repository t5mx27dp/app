package app

import (
	"os"
	"time"
)

type Option func(*App)

func WithSignal(signal os.Signal) Option {
	return func(a *App) {
		a.signals = append(a.signals, signal)

		if a.signal == nil {
			a.signal = make(chan os.Signal, 1)
		}
	}
}

func WithTearDownTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.tearDownTimeout = timeout
	}
}
