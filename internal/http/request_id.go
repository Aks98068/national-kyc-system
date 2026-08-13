package httpserver

import (
	"net/http"

	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

const requestIDHeader = "X-Request-ID"

func requestIDMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			if requestID == "" {
				requestID = generateRequestID()
			}

			ctx := logger.ContextWithRequestID(
				r.Context(),
				requestID,
			)

			w.Header().Set(
				requestIDHeader,
				requestID,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}
