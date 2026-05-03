package bot

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/AlexeyNilov/zero_control/internal/config"
	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TestNewLogsSuccessfulStartup(t *testing.T) {
	originalNewTelegramRunner := newTelegramRunner
	t.Cleanup(func() {
		newTelegramRunner = originalNewTelegramRunner
	})

	newTelegramRunner = func(_ string, _ ...tgbot.Option) (telegramRunner, error) {
		return stubRunner{}, nil
	}

	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)

	_, err := New(config.Config{
		BotToken:    "token",
		PollTimeout: time.Second,
	}, logger, nil)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	assertContains(t, logs.String(), "telegram bot startup successful")
}

func TestDefaultHandlerLogsIncomingMessage(t *testing.T) {
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{logger: logger}

	update := &models.Update{
		ID: 42,
		Message: &models.Message{
			ID:   15,
			Text: "status",
			Chat: models.Chat{
				ID: 777,
			},
			From: &models.User{
				ID:       55,
				Username: "alice",
			},
		},
	}

	telegramBot.handleDefaultUpdate(context.Background(), nil, update)

	logLine := logs.String()
	assertContains(t, logLine, "received telegram message")
	assertContains(t, logLine, "update_id=42")
	assertContains(t, logLine, "message_id=15")
	assertContains(t, logLine, "chat_id=777")
	assertContains(t, logLine, "user_id=55")
	assertContains(t, logLine, "username=alice")
	assertContains(t, logLine, "text=\"status\"")
}

func TestStartHandlerRepliesWithOnlineMessage(t *testing.T) {
	sender := &stubSender{}
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{logger: logger, router: NewRouter(nil)}

	update := &models.Update{
		ID: 7,
		Message: &models.Message{
			ID:   8,
			Text: "/start",
			Chat: models.Chat{
				ID: 999,
			},
			From: &models.User{
				ID:       111,
				Username: "starter",
			},
		},
	}

	err := telegramBot.handleStartCommand(context.Background(), sender, update)
	if err != nil {
		t.Fatalf("handleStartCommand returned error: %v", err)
	}

	if sender.chatID != 999 {
		t.Fatalf("chatID = %d, want %d", sender.chatID, 999)
	}

	if sender.text != "zero_control is online" {
		t.Fatalf("text = %q, want %q", sender.text, "zero_control is online")
	}

	logLine := logs.String()
	assertContains(t, logLine, "received telegram message")
	assertContains(t, logLine, "text=\"/start\"")
}

type stubSender struct {
	chatID int64
	text   string
}

type stubRunner struct{}

func (stubRunner) Start(context.Context) {}

func (s *stubSender) SendMessage(_ context.Context, params *tgbot.SendMessageParams) (*models.Message, error) {
	chatID, ok := params.ChatID.(int64)
	if !ok {
		return nil, tgbot.ErrorForbidden
	}

	s.chatID = chatID
	s.text = params.Text
	return &models.Message{}, nil
}

func assertContains(t *testing.T, value, want string) {
	t.Helper()

	if !strings.Contains(value, want) {
		t.Fatalf("value %q does not contain %q", value, want)
	}
}
