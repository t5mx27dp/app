package app_test

import (
	"context"
	"syscall"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/t5mx27dp/app"
	appmock "github.com/t5mx27dp/app/mock"
)

func TestApp(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		ctx := context.Background()
		runner := appmock.NewRunnable(t)

		runner.On("SetUp", mock.Anything, mock.Anything).Return(nil)
		runner.On("TearDown", mock.Anything).Return()

		app := app.New(ctx, runner, app.WithSignal(syscall.SIGINT), app.WithSignal(syscall.SIGTERM))

		sig := app.GetSignal()

		go func() {
			sig <- syscall.SIGINT
		}()

		err := app.Run()
		require.Nil(t, err)
	})

	t.Run("ContextDone", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runner := appmock.NewRunnable(t)

		runner.On("SetUp", mock.Anything, mock.Anything).Return(nil)
		runner.On("TearDown", mock.Anything).Return()

		app := app.New(ctx, runner, app.WithSignal(syscall.SIGINT), app.WithSignal(syscall.SIGTERM))

		go func() {
			cancel()
		}()

		err := app.Run()
		require.Nil(t, err)
	})

	t.Run("WithoutSignal", func(t *testing.T) {
		ctx := context.Background()
		runner := appmock.NewRunnable(t)

		runner.On("SetUp", mock.Anything, mock.Anything).Return(nil)
		runner.On("TearDown", mock.Anything).Return()

		app := app.New(ctx, runner)

		err := app.Run()
		require.Nil(t, err)
	})
}
