package logger

import "log/slog"

const redactedValue = "[REDACTED]"

func Redact(value any) any {
	return redactedValue
}

func Sensitive(key string) slog.Attr {
	return slog.String(key, redactedValue)
}
