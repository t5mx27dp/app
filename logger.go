package app

import "context"

type Logger interface {
	Error(ctx context.Context, message string, fields map[string]any)
	Warn(ctx context.Context, message string, fields map[string]any)
	Info(ctx context.Context, message string, fields map[string]any)
}
