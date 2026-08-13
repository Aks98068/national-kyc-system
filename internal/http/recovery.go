package httpserver

import (
	"net/http"
	"runtime/debug"

	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

func recoveryMiddleware(
	log *logger.Logger,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					log.Error(
						"HTTP handler panic",
						"request_id",
						logger.RequestIDFromContext(
							r.Context(),
						),
						"panic",
						recovered,
						"stack",
						string(debug.Stack()),
					)

					writeJSON(
						w,
						http.StatusInternalServerError,
						map[string]any{
							"error": map[string]string{
								"code":    "INTERNAL_ERROR",
								"message": "An internal error occurred.",
								"request_id": logger.RequestIDFromContext(
									r.Context(),
								),
							},
						},
					)
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}
