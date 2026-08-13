package logger

import (
	"context"
	"log/slog"
	"os"
)

type Config struct {
	Level  string
	Format string
}

type Logger struct {
	*slog.Logger
}

func New(cfg Config) *Logger {
	level := parseLevel(cfg.Level)

	var handler slog.Handler

	switch cfg.Format {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	return &Logger{
		Logger: slog.New(handler),
	}
}

func parseLevel(value string) slog.Level {
	switch value {
	case "debug":
		return slog.LevelDebug

	case "warn":
		return slog.LevelWarn

	case "error":
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}

func (l *Logger) WithContext(
	ctx context.Context,
) *slog.Logger {
	return l.Logger.With(
		"request_id",
		RequestIDFromContext(ctx),
	)
}
