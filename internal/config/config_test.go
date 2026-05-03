package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexeyNilov/zero_control/internal/config"
)

func TestLoadReadsBotTokenFromDotEnvFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\nDEVELOPER_CHAT_ID=12345\nAUTHORIZED_IDS=11, 22\n")

	cfg, err := config.Load(dir)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.BotToken != "test-token" {
		t.Fatalf("BotToken = %q, want %q", cfg.BotToken, "test-token")
	}

	if cfg.DeveloperChatID != 12345 {
		t.Fatalf("DeveloperChatID = %d, want %d", cfg.DeveloperChatID, 12345)
	}

	if !containsAuthorizedID(cfg.AuthorizedIDs, 11) {
		t.Fatal("AuthorizedIDs does not contain 11")
	}

	if !containsAuthorizedID(cfg.AuthorizedIDs, 22) {
		t.Fatal("AuthorizedIDs does not contain 22")
	}
}

func TestLoadReturnsErrorWhenBotTokenIsMissing(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "DEVELOPER_CHAT_ID=12345\n")

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want missing token error")
	}
}

func TestLoadReturnsErrorWhenDeveloperChatIDIsMissing(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\nAUTHORIZED_IDS=11\n")

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want missing developer chat id error")
	}
}

func TestLoadReturnsErrorWhenDeveloperChatIDIsInvalid(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\nDEVELOPER_CHAT_ID=abc\nAUTHORIZED_IDS=11\n")

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want invalid developer chat id error")
	}
}

func TestLoadReturnsErrorWhenAuthorizedIDsAreMissing(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\nDEVELOPER_CHAT_ID=12345\n")

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want missing authorized ids error")
	}
}

func TestLoadReturnsErrorWhenAuthorizedIDsContainInvalidValue(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\nDEVELOPER_CHAT_ID=12345\nAUTHORIZED_IDS=11,nope\n")

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want invalid authorized ids error")
	}
}

func writeFile(t *testing.T, path, contents string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
}

func containsAuthorizedID(ids map[int64]struct{}, want int64) bool {
	_, ok := ids[want]
	return ok
}
