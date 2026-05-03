package bot

import (
	"bytes"
	"context"
	"errors"
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
		return &stubRunner{}, nil
	}

	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)

	_, err := New(config.Config{
		BotToken:      "token",
		AuthorizedIDs: map[int64]struct{}{55: {}},
		PollTimeout:   time.Second,
	}, logger, nil)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	assertContains(t, logs.String(), "telegram bot startup successful")
}

func TestRunSendsStartupNotificationToDeveloperChat(t *testing.T) {
	runner := &stubRunner{}
	telegramBot := &Bot{
		logger:          log.New(&bytes.Buffer{}, "", 0),
		router:          NewRouter(nil),
		api:             runner,
		developerChatID: 12345,
	}

	err := telegramBot.Run(context.Background())
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if runner.started != 1 {
		t.Fatalf("Start calls = %d, want %d", runner.started, 1)
	}

	if runner.chatID != 12345 {
		t.Fatalf("chatID = %d, want %d", runner.chatID, 12345)
	}

	if runner.text != "zero_control is online" {
		t.Fatalf("text = %q, want %q", runner.text, "zero_control is online")
	}
}

func TestRunReturnsErrorWhenStartupNotificationCannotBeSent(t *testing.T) {
	runner := &stubRunner{sendErr: errors.New("send failed")}
	telegramBot := &Bot{
		logger:          log.New(&bytes.Buffer{}, "", 0),
		router:          NewRouter(nil),
		api:             runner,
		developerChatID: 12345,
	}

	err := telegramBot.Run(context.Background())
	if err == nil {
		t.Fatal("Run returned nil error, want startup notification error")
	}

	assertContains(t, err.Error(), "send startup notification")

	if runner.started != 0 {
		t.Fatalf("Start calls = %d, want %d", runner.started, 0)
	}
}

func TestDefaultHandlerLogsIncomingMessage(t *testing.T) {
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{
		logger:        logger,
		authorizedIDs: map[int64]struct{}{55: {}},
	}

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
	assertNotContains(t, logLine, "text=")
	assertNotContains(t, logLine, "status")
}

func TestDefaultHandlerIgnoresUnauthorizedMessage(t *testing.T) {
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{
		logger:        logger,
		authorizedIDs: map[int64]struct{}{99: {}},
	}

	update := &models.Update{
		ID: 42,
		Message: &models.Message{
			ID:   15,
			Chat: models.Chat{ID: 777},
			From: &models.User{
				ID:       55,
				Username: "alice",
			},
		},
	}

	telegramBot.handleDefaultUpdate(context.Background(), nil, update)

	if logs.Len() != 0 {
		t.Fatalf("logs = %q, want empty log output", logs.String())
	}
}

func TestStatusHandlerRepliesWithOnlineMessage(t *testing.T) {
	sender := &stubSender{}
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{
		logger:        logger,
		router:        NewRouter(nil),
		authorizedIDs: map[int64]struct{}{111: {}},
	}

	update := &models.Update{
		ID: 7,
		Message: &models.Message{
			ID:   8,
			Text: "/status",
			Chat: models.Chat{
				ID: 999,
			},
			From: &models.User{
				ID:       111,
				Username: "starter",
			},
		},
	}

	err := telegramBot.handleStatusCommand(context.Background(), sender, update)
	if err != nil {
		t.Fatalf("handleStatusCommand returned error: %v", err)
	}

	if sender.chatID != 999 {
		t.Fatalf("chatID = %d, want %d", sender.chatID, 999)
	}

	if sender.text != "zero_control is online" {
		t.Fatalf("text = %q, want %q", sender.text, "zero_control is online")
	}

	logLine := logs.String()
	assertContains(t, logLine, "received telegram message")
	assertNotContains(t, logLine, "text=")
	assertNotContains(t, logLine, "/status")
}

func TestStatusHandlerIgnoresUnauthorizedMessage(t *testing.T) {
	sender := &stubSender{}
	var logs bytes.Buffer
	logger := log.New(&logs, "", 0)
	telegramBot := &Bot{
		logger:        logger,
		router:        NewRouter(nil),
		authorizedIDs: map[int64]struct{}{222: {}},
	}

	update := &models.Update{
		ID: 7,
		Message: &models.Message{
			ID:   8,
			Text: "/status",
			Chat: models.Chat{ID: 999},
			From: &models.User{
				ID:       111,
				Username: "starter",
			},
		},
	}

	err := telegramBot.handleStatusCommand(context.Background(), sender, update)
	if err != nil {
		t.Fatalf("handleStatusCommand returned error: %v", err)
	}

	if sender.text != "" {
		t.Fatalf("text = %q, want empty reply", sender.text)
	}

	if logs.Len() != 0 {
		t.Fatalf("logs = %q, want empty log output", logs.String())
	}
}

type stubSender struct {
	chatID int64
	text   string
}

type stubRunner struct {
	started int
	chatID  int64
	text    string
	sendErr error
}

func (s *stubRunner) Start(context.Context) {
	s.started++
}

func (s *stubRunner) SendMessage(_ context.Context, params *tgbot.SendMessageParams) (*models.Message, error) {
	if s.sendErr != nil {
		return nil, s.sendErr
	}

	chatID, ok := params.ChatID.(int64)
	if !ok {
		return nil, tgbot.ErrorForbidden
	}

	s.chatID = chatID
	s.text = params.Text
	return &models.Message{}, nil
}

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

func assertNotContains(t *testing.T, value, unwanted string) {
	t.Helper()

	if strings.Contains(value, unwanted) {
		t.Fatalf("value %q unexpectedly contains %q", value, unwanted)
	}
}
