package httpserver

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /health/live",
		livenessHandler,
	)

	mux.HandleFunc(
		"GET /health/ready",
		readinessHandler,
	)

	return mux
}
