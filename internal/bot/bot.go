package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AlexeyNilov/zero_control/internal/config"
	"github.com/AlexeyNilov/zero_control/internal/service"
	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	logger        *log.Logger
	router        Router
	api           telegramRunner
	startupChatID int64
}

type telegramRunner interface {
	Start(context.Context)
	SendMessage(context.Context, *tgbot.SendMessageParams) (*models.Message, error)
}

var newTelegramRunner = func(token string, options ...tgbot.Option) (telegramRunner, error) {
	return tgbot.New(token, options...)
}

func New(cfg config.Config, logger *log.Logger, control *service.ControlService) (*Bot, error) {
	router := NewRouter(control)
	httpClient := &http.Client{Timeout: cfg.PollTimeout}

	client, err := newTelegramRunner(
		cfg.BotToken,
		tgbot.WithDefaultHandler(func(ctx context.Context, api *tgbot.Bot, update *models.Update) {
			bot := &Bot{logger: logger, router: router, api: api}
			bot.handleDefaultUpdate(ctx, api, update)
		}),
		tgbot.WithMessageTextHandler("/start", tgbot.MatchTypeExact, func(ctx context.Context, api *tgbot.Bot, update *models.Update) {
			bot := &Bot{logger: logger, router: router, api: api}
			if err := bot.handleStartCommand(ctx, api, update); err != nil {
				bot.logger.Printf("start handler error: %v", err)
			}
		}),
		tgbot.WithErrorsHandler(func(err error) {
			logger.Printf("telegram bot error: %v", err)
		}),
		tgbot.WithHTTPClient(cfg.PollTimeout, httpClient),
	)
	if err != nil {
		return nil, fmt.Errorf("initialize telegram bot: %w", err)
	}

	logger.Printf("telegram bot startup successful")

	return &Bot{
		logger:        logger,
		router:        router,
		api:           client,
		startupChatID: cfg.DeveloperChatID,
	}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	if err := b.sendStartupGreeting(ctx); err != nil {
		return err
	}

	b.api.Start(ctx)
	return nil
}

func (b *Bot) sendStartupGreeting(ctx context.Context) error {
	_, err := b.api.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: b.startupChatID,
		Text:   b.router.StartMessage(),
	})
	if err != nil {
		return fmt.Errorf("send startup greeting to chat %s: %w", strconv.FormatInt(b.startupChatID, 10), err)
	}

	return nil
}

func (b *Bot) handleDefaultUpdate(_ context.Context, _ *tgbot.Bot, update *models.Update) {
	message := updateMessage(update)
	if message == nil {
		return
	}

	b.logger.Printf(
		"received telegram message update_id=%d message_id=%d chat_id=%d user_id=%s username=%s",
		update.ID,
		message.ID,
		message.Chat.ID,
		userID(message.From),
		username(message.From),
	)
}

func (b *Bot) handleStartCommand(ctx context.Context, sender messageSender, update *models.Update) error {
	b.handleDefaultUpdate(ctx, nil, update)
	return handleStart(ctx, sender, b.router, update)
}

func updateMessage(update *models.Update) *models.Message {
	if update == nil {
		return nil
	}

	return update.Message
}

func userID(user *models.User) string {
	if user == nil {
		return "unknown"
	}

	return strconv.FormatInt(user.ID, 10)
}

func username(user *models.User) string {
	if user == nil || user.Username == "" {
		return "-"
	}

	return user.Username
}
