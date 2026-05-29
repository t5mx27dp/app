package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

type App interface {
	SetUp(ctx context.Context, spawn Spawn) error
	TearDown()
}

type Spawn func(fn func(ctx context.Context))

type Controller struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	signal  chan os.Signal
	signals []os.Signal

	app App
}

func NewController(app App, opts ...Option) *Controller {
	c := &Controller{
		app: app,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Controller) Run(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)

	c.wg = &sync.WaitGroup{}

	defer func() {
		c.cancel()
		c.wg.Wait()
		c.app.TearDown()
	}()

	err := c.app.SetUp(c.ctx, c.spawn)
	if err != nil {
		return err
	}

	if c.signal != nil {
		signal.Notify(c.signal, c.signals...)
		defer signal.Stop(c.signal)

		select {
		case <-c.ctx.Done():
		case <-c.signal:
		}
	}

	return nil
}

func (c *Controller) spawn(fn func(ctx context.Context)) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		fn(c.ctx)
	}()
}
