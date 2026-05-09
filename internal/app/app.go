package app

import (
	"context"
	"log"

	"github.com/AlexeyNilov/zero_control/internal/bot"
	"github.com/AlexeyNilov/zero_control/internal/config"
	"github.com/AlexeyNilov/zero_control/internal/device"
	"github.com/AlexeyNilov/zero_control/internal/mqtt"
	"github.com/AlexeyNilov/zero_control/internal/service"
)

type App struct {
	bot            *bot.Bot
	mqttSubscriber *mqtt.Subscriber
}

func New(cfg config.Config, logger *log.Logger) (*App, error) {
	deviceService := device.New()
	controlService := service.New(deviceService)
	telegramBot, err := bot.New(cfg, logger, controlService)
	if err != nil {
		return nil, err
	}

	mqttClient := mqtt.NewPahoClient(cfg.MQTTBrokerURL, "zero-control-bot", logger)
	mqttSubscriber := mqtt.NewSubscriber(mqttClient, telegramBot, logger)

	return &App{
		bot:            telegramBot,
		mqttSubscriber: mqttSubscriber,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.mqttSubscriber.Start(ctx); err != nil {
		return err
	}

	return a.bot.Run(ctx)
}
