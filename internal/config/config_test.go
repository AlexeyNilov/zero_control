package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexeyNilov/zero_control/internal/config"
)

func TestLoadReadsBotTokenFromDotEnvFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".env"), "BOT_TOKEN=test-token\n")

	cfg, err := config.Load(dir)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.BotToken != "test-token" {
		t.Fatalf("BotToken = %q, want %q", cfg.BotToken, "test-token")
	}
}

func TestLoadReturnsErrorWhenBotTokenIsMissing(t *testing.T) {
	dir := t.TempDir()

	_, err := config.Load(dir)
	if err == nil {
		t.Fatal("Load returned nil error, want missing token error")
	}
}

func writeFile(t *testing.T, path, contents string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
}
