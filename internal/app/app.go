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

func New(cfg config.Config, logger *log.Logger) *App {
	deviceService := device.New()
	controlService := service.New(deviceService)
	telegramBot := bot.New(cfg, logger, controlService)

	return &App{bot: telegramBot}
}

func (a *App) Run(ctx context.Context) error {
	return a.bot.Run(ctx)
}
