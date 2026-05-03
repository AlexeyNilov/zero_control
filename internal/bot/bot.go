package bot

import (
	"context"
	"log"

	"github.com/AlexeyNilov/zero_control/internal/config"
	"github.com/AlexeyNilov/zero_control/internal/service"
)

type Bot struct {
	logger  *log.Logger
	router  Router
	timeout string
}

func New(cfg config.Config, logger *log.Logger, control *service.ControlService) *Bot {
	return &Bot{
		logger:  logger,
		router:  NewRouter(control),
		timeout: cfg.PollTimeout.String(),
	}
}

func (b *Bot) Run(context.Context) error {
	b.logger.Printf("bot skeleton initialized with poll timeout %s", b.timeout)
	return nil
}
