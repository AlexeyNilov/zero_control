package main

import (
	"context"
	"os"

	"github.com/AlexeyNilov/zero_control/internal/app"
	"github.com/AlexeyNilov/zero_control/internal/config"
	logging "github.com/AlexeyNilov/zero_control/internal/logging"
)

func main() {
	logger := logging.New(os.Stdout)

	cfg, err := config.Load(".")
	if err != nil {
		logger.Printf("config error: %v", err)
		os.Exit(1)
	}

	application := app.New(cfg, logger)
	if err := application.Run(context.Background()); err != nil {
		logger.Printf("application error: %v", err)
		os.Exit(1)
	}
}
