package app

import (
	"net/http"

	"github.com/national-kyc-system/national-kyc-system/internal/config"
	httpserver "github.com/national-kyc-system/national-kyc-system/internal/http"
	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

type App struct {
	Config config.Config
	Logger *logger.Logger
	Server *httpserver.Server
}

func New(cfg config.Config) *App {
	log := logger.New(
		logger.Config{
			Level:  cfg.Logging.Level,
			Format: cfg.Logging.Format,
		},
	)

	router := httpserver.NewRouter()

	handler := applyHTTPMiddleware(
		cfg,
		log,
		router,
	)

	server := httpserver.New(
		cfg.HTTP,
		log,
		handler,
	)

	return &App{
		Config: cfg,
		Logger: log,
		Server: server,
	}
}

func applyHTTPMiddleware(
	cfg config.Config,
	log *logger.Logger,
	handler http.Handler,
) http.Handler {
	return httpserver.ApplyMiddleware(
		cfg.HTTP,
		log,
		handler,
	)
}
