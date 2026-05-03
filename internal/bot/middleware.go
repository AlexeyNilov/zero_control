package bot

import "context"

type Handler func(context.Context) string

func LoggingMiddleware(next Handler) Handler {
	return func(ctx context.Context) string {
		return next(ctx)
	}
}
