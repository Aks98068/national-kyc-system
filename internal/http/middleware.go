package httpserver

import (
	"net/http"

	"github.com/national-kyc-system/national-kyc-system/internal/config"
	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

func ApplyMiddleware(
	cfg config.HTTPConfig,
	log *logger.Logger,
	handler http.Handler,
) http.Handler {
	handler = bodyLimitMiddleware(
		cfg,
		handler,
	)

	handler = securityHeadersMiddleware(
		handler,
	)

	handler = recoveryMiddleware(
		log,
		handler,
	)

	handler = requestIDMiddleware(
		handler,
	)

	return handler
}
