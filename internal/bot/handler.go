package bot

import (
	"context"
	"fmt"
	"strconv"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type messageSender interface {
	SendMessage(context.Context, *tgbot.SendMessageParams) (*models.Message, error)
}

func handleStatus(ctx context.Context, sender messageSender, router Router, update *models.Update) error {
	if update == nil || update.Message == nil {
		return nil
	}

	_, err := sender.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   router.StatusMessage(),
	})
	if err != nil {
		return fmt.Errorf("send /status reply to chat %s: %w", strconv.FormatInt(update.Message.Chat.ID, 10), err)
	}

	return nil
}
