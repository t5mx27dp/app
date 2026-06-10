package app

import "context"

type Logger interface {
	Log(ctx context.Context, message string, fields map[string]any)
	Error(ctx context.Context, err error, fields map[string]any)
}
