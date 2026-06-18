package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Runnable interface {
	SetUp(ctx context.Context, spawn Spawn) error
	TearDown(ctx context.Context)
}

type Spawn func(fn func(ctx context.Context))

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	signal  chan os.Signal
	signals []os.Signal

	tearDownTimeout time.Duration

	runner Runnable
}

func New(ctx context.Context, runner Runnable, opts ...Option) *App {
	ctx, cancel := context.WithCancel(ctx)

	a := &App{
		ctx:    ctx,
		cancel: cancel,
		runner: runner,
	}

	for _, opt := range opts {
		opt(a)
	}

	if a.tearDownTimeout <= 0 {
		a.tearDownTimeout = time.Second * 10
	}

	return a
}

func (a *App) Run() error {
	defer func() {
		a.cancel()
		a.tearDown()
		a.wg.Wait()
	}()

	err := a.setUp()
	if err != nil {
		return err
	}

	a.listen()

	return nil
}

func (a *App) setUp() error {
	return a.runner.SetUp(a.ctx, a.spawn)
}

func (a *App) tearDown() {
	ctx, cancel := context.WithTimeout(
		context.WithoutCancel(a.ctx),
		a.tearDownTimeout,
	)
	defer cancel()

	a.runner.TearDown(ctx)
}

func (a *App) listen() {
	if a.signal == nil {
		return
	}

	signal.Notify(a.signal, a.signals...)
	defer signal.Stop(a.signal)

	select {
	case <-a.ctx.Done():
	case <-a.signal:
	}
}

func (a *App) spawn(fn func(ctx context.Context)) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		fn(a.ctx)
	}()
}
