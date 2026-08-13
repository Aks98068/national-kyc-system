package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/national-kyc-system/national-kyc-system/internal/config"
	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

type Server struct {
	httpServer *http.Server
	logger     *logger.Logger
	config     config.HTTPConfig
}

func New(
	cfg config.HTTPConfig,
	log *logger.Logger,
	handler http.Handler,
) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:           handler,
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			MaxHeaderBytes:    1 << 20,
		},
		logger: log,
		config: cfg,
	}
}

func (s *Server) Start() error {
	s.logger.Info(
		"HTTP server starting",
		"address",
		s.httpServer.Addr,
	)

	err := s.httpServer.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.config.ShutdownTimeout,
	)
	defer cancel()

	s.logger.Info("HTTP server shutting down")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf(
			"HTTP server shutdown failed: %w",
			err,
		)
	}

	s.logger.Info("HTTP server stopped")

	return nil
}

func DefaultTimeouts() (
	read time.Duration,
	header time.Duration,
	write time.Duration,
	idle time.Duration,
) {
	read = 15 * time.Second
	header = 5 * time.Second
	write = 30 * time.Second
	idle = 60 * time.Second

	return
}
