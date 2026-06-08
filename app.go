package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

type Runnable interface {
	SetUp(ctx context.Context, spawn Spawn) error
	TearDown()
}

type Spawn func(fn func(ctx context.Context))

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	signal  chan os.Signal
	signals []os.Signal

	runner Runnable
}

func New(runner Runnable, opts ...Option) *App {
	a := &App{
		runner: runner,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *App) Run(ctx context.Context) error {
	a.ctx, a.cancel = context.WithCancel(ctx)

	a.wg = &sync.WaitGroup{}

	defer func() {
		a.cancel()
		a.wg.Wait()
		a.runner.TearDown()
	}()

	err := a.runner.SetUp(a.ctx, a.spawn)
	if err != nil {
		return err
	}

	if a.signal != nil {
		signal.Notify(a.signal, a.signals...)
		defer signal.Stop(a.signal)

		select {
		case <-a.ctx.Done():
		case <-a.signal:
		}
	}

	return nil
}

func (a *App) spawn(fn func(ctx context.Context)) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		fn(a.ctx)
	}()
}
