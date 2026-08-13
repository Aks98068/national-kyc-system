package logger

import "context"

type contextKey string

const requestIDKey contextKey = "request_id"

func ContextWithRequestID(
	ctx context.Context,
	requestID string,
) context.Context {
	return context.WithValue(
		ctx,
		requestIDKey,
		requestID,
	)
}

func RequestIDFromContext(
	ctx context.Context,
) string {
	value := ctx.Value(requestIDKey)

	if value == nil {
		return ""
	}

	requestID, ok := value.(string)
	if !ok {
		return ""
	}

	return requestID
}
