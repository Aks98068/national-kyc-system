package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/national-kyc-system/national-kyc-system/internal/app"
	"github.com/national-kyc-system/national-kyc-system/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"configuration error: %v\n",
			err,
		)

		os.Exit(1)
	}

	application := app.New(cfg)

	go func() {
		if err := application.Server.Start(); err != nil {
			application.Logger.Error(
				"HTTP server stopped unexpectedly",
				"error",
				err,
			)

			os.Exit(1)
		}
	}()

	signalChannel := make(
		chan os.Signal,
		1,
	)

	signal.Notify(
		signalChannel,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-signalChannel

	if err := application.Server.Shutdown(); err != nil {
		application.Logger.Error(
			"graceful shutdown failed",
			"error",
			err,
		)

		os.Exit(1)
	}
}
