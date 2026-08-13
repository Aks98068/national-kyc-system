package httpserver

import "net/http"

func securityHeadersMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(
				"X-Content-Type-Options",
				"nosniff",
			)

			w.Header().Set(
				"X-Frame-Options",
				"DENY",
			)

			w.Header().Set(
				"Referrer-Policy",
				"no-referrer",
			)

			w.Header().Set(
				"Cache-Control",
				"no-store",
			)

			w.Header().Set(
				"Permissions-Policy",
				"camera=(), microphone=(), geolocation=()",
			)

			next.ServeHTTP(w, r)
		},
	)
}
