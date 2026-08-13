package logger

import (
	"context"
	"log/slog"
	"testing"
)

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected slog.Level
	}{
		{
			name:     "debug",
			input:    "debug",
			expected: slog.LevelDebug,
		},
		{
			name:     "info",
			input:    "info",
			expected: slog.LevelInfo,
		},
		{
			name:     "warn",
			input:    "warn",
			expected: slog.LevelWarn,
		},
		{
			name:     "error",
			input:    "error",
			expected: slog.LevelError,
		},
		{
			name:     "unknown defaults to info",
			input:    "something",
			expected: slog.LevelInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLevel(tt.input)

			if got != tt.expected {
				t.Fatalf(
					"expected %v, got %v",
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestRequestIDContext(t *testing.T) {
	ctx := context.Background()

	ctx = ContextWithRequestID(
		ctx,
		"test-request-id",
	)

	got := RequestIDFromContext(ctx)

	if got != "test-request-id" {
		t.Fatalf(
			"expected request ID %q, got %q",
			"test-request-id",
			got,
		)
	}
}

func TestMissingRequestID(t *testing.T) {
	ctx := context.Background()

	got := RequestIDFromContext(ctx)

	if got != "" {
		t.Fatalf(
			"expected empty request ID, got %q",
			got,
		)
	}
}

func TestRedact(t *testing.T) {
	got := Redact("sensitive-value")

	if got != redactedValue {
		t.Fatalf(
			"expected %q, got %v",
			redactedValue,
			got,
		)
	}
}
