package bot

import (
	"context"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler func(context.Context, *tgbot.Bot, *models.Update)

func LoggingMiddleware(next Handler) Handler {
	return func(ctx context.Context, bot *tgbot.Bot, update *models.Update) {
		next(ctx, bot, update)
	}
}
