package app

import (
	"context"
	"log"

	"github.com/AlexeyNilov/zero_control/internal/bot"
	"github.com/AlexeyNilov/zero_control/internal/config"
	"github.com/AlexeyNilov/zero_control/internal/device"
	"github.com/AlexeyNilov/zero_control/internal/service"
)

type App struct {
	bot *bot.Bot
}

func New(cfg config.Config, logger *log.Logger) (*App, error) {
	deviceService := device.New()
	controlService := service.New(deviceService)
	telegramBot, err := bot.New(cfg, logger, controlService)
	if err != nil {
		return nil, err
	}

	return &App{bot: telegramBot}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.bot.Run(ctx)
}
