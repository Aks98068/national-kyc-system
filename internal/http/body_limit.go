package httpserver

import (
	"net/http"

	"github.com/national-kyc-system/national-kyc-system/internal/config"
)

func bodyLimitMiddleware(
	cfg config.HTTPConfig,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(
				w,
				r.Body,
				cfg.MaxBodyBytes,
			)

			next.ServeHTTP(w, r)
		},
	)
}
