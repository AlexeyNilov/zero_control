package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const defaultPollTimeout = 30 * time.Second
const defaultMQTTBrokerURL = "tcp://localhost:1883"

type Config struct {
	BotToken        string
	DeveloperChatID int64
	MainChatID      int64
	AuthorizedIDs   map[int64]struct{}
	PollTimeout     time.Duration
	MQTTBrokerURL   string
}

func Load(dir string) (Config, error) {
	values, err := readDotEnv(filepath.Join(dir, ".env"))
	if err != nil {
		return Config{}, err
	}

	token := firstNonEmpty(os.Getenv("BOT_TOKEN"), values["BOT_TOKEN"])
	if token == "" {
		return Config{}, errors.New("BOT_TOKEN is required")
	}

	developerChatID, err := parseChatID("DEVELOPER_CHAT_ID", firstNonEmpty(os.Getenv("DEVELOPER_CHAT_ID"), values["DEVELOPER_CHAT_ID"]))
	if err != nil {
		return Config{}, err
	}

	mainChatID, err := parseChatID("MAIN_CHAT_ID", firstNonEmpty(os.Getenv("MAIN_CHAT_ID"), values["MAIN_CHAT_ID"]))
	if err != nil {
		return Config{}, err
	}

	authorizedIDs, err := parseAuthorizedIDs(firstNonEmpty(os.Getenv("AUTHORIZED_IDS"), values["AUTHORIZED_IDS"]))
	if err != nil {
		return Config{}, err
	}

	return Config{
		BotToken:        token,
		DeveloperChatID: developerChatID,
		MainChatID:      mainChatID,
		AuthorizedIDs:   authorizedIDs,
		PollTimeout:     defaultPollTimeout,
		MQTTBrokerURL:   firstNonEmpty(os.Getenv("MQTT_BROKER_URL"), values["MQTT_BROKER_URL"], defaultMQTTBrokerURL),
	}, nil
}

func parseChatID(name, value string) (int64, error) {
	if value == "" {
		return 0, fmt.Errorf("%s is required", name)
	}

	chatID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer: %w", name, err)
	}

	return chatID, nil
}

func parseAuthorizedIDs(value string) (map[int64]struct{}, error) {
	if value == "" {
		return nil, errors.New("AUTHORIZED_IDS is required")
	}

	ids := make(map[int64]struct{})
	for _, field := range strings.Split(value, ",") {
		id, err := parseAuthorizedID(field)
		if err != nil {
			return nil, err
		}

		ids[id] = struct{}{}
	}

	return ids, nil
}

func parseAuthorizedID(value string) (int64, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return 0, errors.New("AUTHORIZED_IDS must contain only valid integers")
	}

	id, err := strconv.ParseInt(trimmed, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("AUTHORIZED_IDS must contain only valid integers: %w", err)
	}

	return id, nil
}

func readDotEnv(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return map[string]string{}, nil
		}
		return nil, fmt.Errorf("read .env: %w", err)
	}

	return parseEnv(string(data))
}

func parseEnv(data string) (map[string]string, error) {
	values := make(map[string]string)

	for lineNumber, line := range strings.Split(data, "\n") {
		if err := parseLine(values, line, lineNumber+1); err != nil {
			return nil, err
		}
	}

	return values, nil
}

func parseLine(values map[string]string, raw string, lineNumber int) error {
	line := strings.TrimSpace(raw)
	if line == "" || strings.HasPrefix(line, "#") {
		return nil
	}

	key, value, ok := strings.Cut(line, "=")
	if !ok {
		return fmt.Errorf(".env line %d: invalid format", lineNumber)
	}

	values[strings.TrimSpace(key)] = strings.TrimSpace(value)
	return nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
